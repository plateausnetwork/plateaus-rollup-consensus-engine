package dicebear

import (
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/rhizomplatform/plateaus-rollup-consensus-engine/internal/lottery/nft"
	httpDicebear "github.com/rhizomplatform/plateaus-rollup-consensus-engine/pkg/dicebear/http"
	"github.com/rhizomplatform/plateaus-rollup-consensus-engine/pkg/http"
	"github.com/stretchr/testify/assert"
	"io"
	http2 "net/http"
	"strings"
	"testing"
)

func TestService_Generate(t *testing.T) {
	hash := "hash"
	ctrl := gomock.NewController(t)

	clientMock := http.NewMockClientDoer(ctrl)
	c := httpDicebear.NewClient(clientMock)

	clientRequest, _ := http2.NewRequest(http2.MethodGet, fmt.Sprintf("%s/%s/%s.svg", c.Url, "bottts", hash), nil)
	clientResponse := &http2.Response{
		Body: io.NopCloser(strings.NewReader("<svg>some svg</svg>")),
	}
	clientMock.EXPECT().Do(gomock.Eq(clientRequest)).Return(clientResponse, nil)

	s := NewService(c)

	res, err := s.Generate("hash")

	assert.IsType(t, &nft.LotteryValidation{}, res)
	assert.Equal(t, []byte("<svg>some svg</svg>"), res.Image)
	assert.NoError(t, err)
}

func TestService_Generate_GetAvatar_Error(t *testing.T) {
	hash := "hash"
	ctrl := gomock.NewController(t)

	clientMock := http.NewMockClientDoer(ctrl)
	c := httpDicebear.NewClient(clientMock)

	clientRequest, _ := http2.NewRequest(http2.MethodGet, fmt.Sprintf("%s/%s/%s.svg", c.Url, "bottts", hash), nil)
	clientMock.EXPECT().Do(gomock.Eq(clientRequest)).Return(nil, errors.New("some error"))

	s := NewService(c)

	_, err := s.Generate("hash")

	assert.Error(t, err)
}
