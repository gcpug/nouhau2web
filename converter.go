package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type Converter struct {
	gcs *StorageService
}

func NewConverter(gcs *StorageService) *Converter {
	return &Converter{gcs: gcs}
}

func (c *Converter) Run(ctx context.Context, bucket string, fl *FileList) error {
	for _, fn := range fl.CurrentFileList {
		lfp := fmt.Sprintf("%s/%s", fl.Dir, fn)
		if err := c.LocalToGCS(ctx, lfp, bucket, fn); err != nil {
			return err
		}
	}

	for _, ufl := range fl.UnderFileList {
		if err := c.Run(ctx, bucket, ufl); err != nil {
			return err
		}
	}
	return nil
}

func (c *Converter) LocalToGCS(ctx context.Context, localFilePath string, bucket string, object string) (rerr error) {
	f, err := os.Open(localFilePath)
	if err != nil {
		return err
	}

	gcsw := c.gcs.NewWriter(ctx, bucket, object)
	defer func() {
		if err := gcsw.Close(); err != nil {
			rerr = err
		}
	}()

	gcsw.ObjectAttrs.ContentType = c.ContentType(object)
	gcsw.ObjectAttrs.ContentDisposition = fmt.Sprintf(`attachment;filename="%v"`, object)
	_, err = io.Copy(gcsw, f)
	if err != nil {
		return err
	}

	return nil
}

// Extension is 拡張子からContentTypeを指定する
func (c *Converter) ContentType(fileName string) string {
	ext := filepath.Ext(fileName)

	ct := ""
	switch ext {
	default:
		ct = "application/octet-stream"
	case ".html", ".htm":
		ct = "text/html;charset=utf-8"
	case ".css":
		ct = "text/css;charset=utf-8"
	case ".js":
		ct = "text/javascript;charset=utf-8"
	case ".jpeg", ".jpg":
		ct = "image/jpeg"
	case ".png":
		ct = "image/png"
	case ".gif":
		ct = "image/gif"
	case ".txt":
		ct = "text/plain;charset=utf-8"
	case ".json":
		ct = "application/json;charset=utf-8"
	case ".pdf":
		ct = "application/pdf"
	case ".ico":
		ct = "image/x-icon"
	}

	return ct
}
