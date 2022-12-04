package tohttps

import (
	"net/http"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
)

// Well, I really didn't want to do this, but some websites just blocked my
// plain request, resulting in failure checking if they support HTTPS.
const (
	ua          = `Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36`
	timeoutSecs = 10

	errNameResolution   = "Temporary failure in name resolution"
	errCertificate      = "x509"
	errRefused          = "connection refused"
	errReset            = "connection reset by peer"
	errInternal         = "remote error: tls: internal error"
	errHandshake        = "handshake"
	errHttpResponse     = "HTTP response to HTTPS"
	errEOF              = "EOF"
	errTooManyRedirects = "stopped after 10 redirects"
	errUnrecognizedName = "remote error: tls: unrecognized name"
	errIoTimeout        = "i/o timeout"
	errUnsupported      = "unsupported protocol version"
	errNoAssociated     = "No address associated with hostname"
	errProtocol         = "PROTOCOL_ERROR"
	errNoSuchHost       = "no such host"
	errNoRoute          = "no route to host"
	errConnTimeout      = "connection timed out"
	errMalformedCode    = "malformed HTTP status code"
	errTimeoutExceeded  = "Client.Timeout exceeded"
	errStreamClosed     = "STREAM_CLOSED"
)

func httpsSupported(url string) (bool, error) {
	strippedUrl := stripUrl(url)
	httpsUrl := "https://" + strippedUrl

	req, err := http.NewRequest("GET", httpsUrl, nil)
	if err != nil {
		log.WithFields(log.Fields{
			"url":   httpsUrl,
			"error": err.Error(),
		}).Error("Failed creating new https request")
		return false, nil
	}
	req.Header.Add("User-Agent", ua)

	client := &http.Client{
		Timeout: time.Second * timeoutSecs,
	}
	_, err = client.Do(req)

	if err != nil {
		switch {
		case strings.HasSuffix(err.Error(), errNameResolution):
			log.WithFields(log.Fields{
				"url": httpsUrl,
			}).Debug("Name resolution failed")
		case strings.Contains(err.Error(), errCertificate):
			log.WithFields(log.Fields{
				"url": httpsUrl,
			}).Debug("Got wrong certificate")
		case strings.Contains(err.Error(), errRefused):
			log.WithFields(log.Fields{
				"url": httpsUrl,
			}).Debug("Connection refused")
		case strings.Contains(err.Error(), errReset):
			log.WithFields(log.Fields{
				"url": httpsUrl,
			}).Debug("Connection reset")
		case strings.Contains(err.Error(), errInternal):
			log.WithFields(log.Fields{
				"url": httpsUrl,
			}).Debug("Internal tls error from server")
		case strings.Contains(err.Error(), errHandshake):
			log.WithFields(log.Fields{
				"url": httpsUrl,
			}).Debug("Failed TLS handshake")
		case strings.Contains(err.Error(), errHttpResponse):
			log.WithFields(log.Fields{
				"url": httpsUrl,
			}).Debug("Server responded with HTTP")
		case strings.Contains(err.Error(), errEOF):
			log.WithFields(log.Fields{
				"url": httpsUrl,
			}).Debug("Server responded EOF")
		case strings.Contains(err.Error(), errTooManyRedirects):
			log.WithFields(log.Fields{
				"url": httpsUrl,
			}).Debug("Request stopped after 10 redirects")
		case strings.Contains(err.Error(), errUnrecognizedName):
			log.WithFields(log.Fields{
				"url": httpsUrl,
			}).Debug("Unrecognized TLS name")
		case strings.Contains(err.Error(), errIoTimeout):
			log.WithFields(log.Fields{
				"url": httpsUrl,
			}).Debug("I/O timeout")
		case strings.Contains(err.Error(), errUnsupported):
			log.WithFields(log.Fields{
				"url": httpsUrl,
			}).Debug("Unsupported TLS version")
		case strings.Contains(err.Error(), errNoAssociated):
			log.WithFields(log.Fields{
				"url": httpsUrl,
			}).Debug("No associated address to url")
		case strings.Contains(err.Error(), errProtocol):
			log.WithFields(log.Fields{
				"url": httpsUrl,
			}).Debug("Protocol error")
		case strings.Contains(err.Error(), errNoSuchHost):
			log.WithFields(log.Fields{
				"url": httpsUrl,
			}).Debug("No such host")
		case strings.Contains(err.Error(), errNoRoute):
			log.WithFields(log.Fields{
				"url": httpsUrl,
			}).Debug("No route to host")
		case strings.Contains(err.Error(), errConnTimeout):
			log.WithFields(log.Fields{
				"url": httpsUrl,
			}).Debug("Connection timed out")
		case strings.Contains(err.Error(), errMalformedCode):
			log.WithFields(log.Fields{
				"url": httpsUrl,
			}).Debug("Got malformed HTTP status code")
		case strings.Contains(err.Error(), errStreamClosed):
			log.WithFields(log.Fields{
				"url": httpsUrl,
			}).Debug("Stream closed")
		case strings.Contains(err.Error(), errTimeoutExceeded):
			log.WithFields(log.Fields{
				"url": httpsUrl,
			}).Debugf("Server not responded in %d seconds", timeoutSecs)
		default:
			log.WithFields(log.Fields{
				"url":   httpsUrl,
				"error": err.Error(),
			}).Debug("Wrong response from server")
		}
		return false, nil
	}

	return true, nil
}
