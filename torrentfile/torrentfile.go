import (
	"github.come/jackpal/bencode-go"
)

type bencodeInfo struct {
	Pieces string 'bencode:"pieces"'
	PiecesLength int 'bencode:"piece length"'
	Length int 'bencode:"length"'
	Name string 'bencode:"name"'
}

type bencode Torrent struct {
	Announce string 'bencode:"announce"'
	Infro bencodeInfo 'bencode:"info"'
}

func open(r io.Reader) (*bencodeTorrent, error) {
	bto := bencodeTorrent{}
	err := bencode.Unmarshal(r, &bto)
	if err != nil {
		return nil, err
	}
	return &bto, nil
}

