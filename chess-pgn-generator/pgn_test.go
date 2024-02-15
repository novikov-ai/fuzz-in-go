package pgn

import (
	"net/url"
	"strings"
	"testing"
)

func FuzzGetPictureURL(f *testing.F) {
	testcases := []string{
		"e4 e6 d4 d5 Nd2 Nf6 Bd3 Be7 Ngf3 c5 dxc5 Bxc5 O-O O-O Re1 Ng4 Rf1 Nc6 h3 Nf6 Qe2 Re8 Rd1 Qb6 Nb3 dxe4 Nxc5 exf3 Qe3 Nd5 Qe4 Qxc5 Qxh7+ Kf8 Qh8+ Ke7 Qxg7 Bd7 Bg5+ Kd6 c4 Rg8 Qh6 f6 cxd5 Rxg5 dxc6",
		"e4 e6 d4 d5 Nd2 Nf6 Bd3 Be7 Ngf3 c5 dxc5 Bxc5 O-O O-O Re1 Ng4 Rf1 Nc6 h3 Nf6 Qe2 Re8 Rd1 Qb6 Nb3 dxe4 Nxc5 exf3 Qe3 Nd5 Qe4 Qxc5 Qxh7+ Kf8 Qh8+ Ke7 Qxg7 Bd7 Bg5+ Kd6 c4 Rg8 Qh6 f6 cxd5 Rxg5 dxc6 Rxg2",
		"e4 e6 d4 d5 Nd2 Nf6 Bd3",
		"e4 e6",
	}

	for _, tc := range testcases {
		f.Add(tc)
	}
	f.Fuzz(func(t *testing.T, orig string) {
		fetchedURL, err := GetPictureURL(orig)
		if err != nil {
			return
		}

		if fetchedURL == "" {
			t.Errorf("url is empty")
		}

		if _, err := url.ParseRequestURI(fetchedURL); err != nil {
			t.Errorf("url is incorect: %s", err)
		}

		if !strings.HasPrefix(fetchedURL, "https://lichess1.org/export/fen.gif?fen=") {
			t.Errorf("lichess url prefix is incorect:%s", fetchedURL)
		}
	})
}
