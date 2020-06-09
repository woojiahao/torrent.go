package p2p

import (
	"github.com/woojiahao/torrent.go/internal/client"
	"github.com/woojiahao/torrent.go/internal/piece"
	"github.com/woojiahao/torrent.go/internal/tracker"
	"log"
	"runtime"
	"time"
)

const maxBlockSize = 16384

type Torrent struct {
	Peers       []tracker.Peer
	InfoHash    string
	PeerID      string
	Pieces      piece.Pieces
	PieceLength int
	Length      int
}

type pieceResult struct {
	index int
	buf   []byte
}

// Download a single piece as mandated by a pieceWork
// Pieces are downloaded sequentially
func downloadPiece(c *client.Client, work *pieceWork) ([]byte, error) {
	log.Print("trying to download piece ", work.index)
	// While downloading the file, we want to keep track of the progress of downloading each block
	// of the piece
	progress := pieceProgress{
		index:  work.index,
		client: c,
		buf:    make([]byte, work.length),
	}

	// A deadline is used to drop the connection if there is no more messages
	_ = c.Conn.Conn.SetDeadline(time.Now().Add(30 * time.Second))
	defer func() {
		// Make sure to reset the deadline for future uses
		_ = c.Conn.Conn.SetDeadline(time.Time{})
	}()

	// As long as all the pieces are not confirmed to have been downloaded, make a download request
	for progress.downloaded < work.length {
		//If the peer isn't choked at the moment, request for a new block
		if !progress.client.Choked {
			blockSize := maxBlockSize

			if blockSize > (work.length - progress.requested) {
				blockSize = work.length - progress.requested
			}

			err := c.SendRequest(work.index, progress.requested, blockSize)
			if err != nil {
				return nil, err
			}

			progress.requested += blockSize
		}

		// In either case, we want to read from the peer and see what kind of message it is emitting
		err := progress.read()
		if err != nil {
			return nil, err
		}
	}

	return progress.buf, nil
}

func (t *Torrent) startPeerDownload(peer tracker.Peer, workQueue chan *pieceWork, results chan *pieceResult) {
	log.Print("starting peer download for ", peer.Address())
	c, err := client.New(peer, t.InfoHash, t.PeerID)
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer func() {
		if c.Conn != nil {
			_ = c.Conn.Conn.Close()
		}
	}()

	_ = c.SendInterested()

	// Every client connection will attempt to take a piece of work to perform
	for work := range workQueue {
		log.Printf("downloading piece %d\n", work.index)
		// If the connected peer does not contain the piece, re-add the piece of work
		// and move on to the next piece of work
		if !c.Bitfield.HasPiece(work.index) {
			workQueue <- work
			continue
		}

		buf, err := downloadPiece(c, work)
		if err != nil {
			log.Print("failed to download piece ", err.Error())
			workQueue <- work
			continue
		}

		err = work.checkIntegrity(buf)
		if err != nil {
			// If the downloaded piece's hash is not the same as the target hash, re-add the piece of work
			// for another connection to try again
			log.Print(err.Error())
			workQueue <- work
			continue
		}

		// Inform the peer that we now have downloaded the piece
		_ = c.SendHave(work.index)
		// Add a new piece result to the results channel
		results <- &pieceResult{work.index, buf}
	}
}

func (t *Torrent) calculatePieceBounds(pos int) (int, int) {
	begin := pos * t.PieceLength
	end := begin + t.PieceLength
	if end > t.Length {
		end = t.Length
	}
	return begin, end
}

// As the file is broken down into different pieces, we need to calculate the size of the piece
// based on its position within the file
func (t *Torrent) calculatePieceSize(pos int) int {
	begin, end := t.calculatePieceBounds(pos)
	return end - begin
}

func (t *Torrent) Download() []byte {
	log.Println("starting torrent download")

	// Work queue to store the remaining pieces left to download
	workQueue := make(chan *pieceWork, len(t.Pieces))

	// Resulting pieces to store
	results := make(chan *pieceResult)

	// Fill the work queue with the pieces to download
	for i, hash := range t.Pieces {
		length := t.calculatePieceSize(i)
		workQueue <- &pieceWork{i, hash, length}
	}

	log.Printf("work queue is %v\n", workQueue)

	// Begin downloading the pieces
	for _, peer := range t.Peers {
		go t.startPeerDownload(peer, workQueue, results)
	}

	// As we begin receiving pieces, we can start to add them to a larger buffer
	buf := make([]byte, t.Length)
	donePieces := 0
	for donePieces < len(t.Pieces) {
		res := <-results
		begin, end := t.calculatePieceBounds(res.index)
		copy(buf[begin:end], res.buf)
		donePieces++

		percent := float64(donePieces) / float64(len(t.Pieces)) * 100
		numWorkers := runtime.NumGoroutine() - 1 // exclude main thread
		log.Printf("(%0.2f%%) Downloaded piece %d from %d peers", percent, res.index, numWorkers)
	}

	close(workQueue)

	return buf
}
