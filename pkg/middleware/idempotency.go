package middleware

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yusufekoanggoro/payment-service/internal/modules/payment/domain"
)

// Menangkap Output
type bodyCaptureWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w *bodyCaptureWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (m *middleware) Idempotency() gin.HandlerFunc {
	return func(c *gin.Context) {
		key := c.GetHeader("Idempotency-Key")
		if key == "" {
			c.Next()
			return
		}

		ctx := c.Request.Context()

		// Baca dan hash request body
		bodyBytes, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.AbortWithStatusJSON(500, gin.H{"error": "failed to read request body"})
			return
		}
		requestHash := sha256.Sum256(bodyBytes)
		hashHex := hex.EncodeToString(requestHash[:])

		// Restore body agar handler bisa membaca ulang
		c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

		// Cek apakah idempotency key sudah ada dan hash cocok
		stored, err := m.repo.GetByKey(ctx, key, hashHex)
		if err != nil {
			c.AbortWithStatusJSON(500, gin.H{"error": "failed to check idempotency"})
			return
		}
		if stored != nil {
			if stored.RequestHash != hashHex {
				c.AbortWithStatusJSON(409, gin.H{"error": "idempotency key reuse with different payload"})
				return
			}
			c.Data(stored.StatusCode, "application/json", []byte(stored.ResponseBody))
			c.Abort()
			return
		}

		// Simpan sementara status processing
		err = m.repo.Save(ctx, &domain.IdempotencyKey{
			Key:          key,
			RequestHash:  hashHex,
			StatusCode:   0,
			ResponseBody: "",
			CreatedAt:    time.Now(),
		})
		if err != nil {
			c.AbortWithStatusJSON(500, gin.H{"error": "failed to initialize idempotency"})
			return
		}

		// Tangkap response hasil handler
		writer := &bodyCaptureWriter{
			ResponseWriter: c.Writer,
			body:           &bytes.Buffer{},
		}
		c.Writer = writer

		c.Next()

		// Save actual response result
		_ = m.repo.Update(ctx, &domain.IdempotencyKey{
			Key:          key,
			StatusCode:   c.Writer.Status(),
			ResponseBody: writer.body.String(),
		})
	}
}
