package main

import (
	"context"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/russross/blackfriday"
)

type Converter struct {
	root string
	gcs  *StorageService
}

func NewConverter(gcs *StorageService, root string) *Converter {
	return &Converter{
		root: root,
		gcs:  gcs,
	}
}

func (c *Converter) Run(ctx context.Context, bucket string, fl *FileList, parent []string) error {
	var current = parent
	current = append(current, fl.Dir)

	for _, fn := range fl.CurrentFileList {
		fp := current
		fp = append(fp, fn)
		lfp := strings.Join(fp, "/")
		if err := c.Process(ctx, lfp, bucket, fn); err != nil {
			return err
		}
	}

	for _, ufl := range fl.UnderFileList {
		if err := c.Run(ctx, bucket, ufl, current); err != nil {
			return err
		}
	}
	return nil
}

func (c *Converter) MarkdownToHTML(md []byte) []byte {
	return blackfriday.Run(md)
}

func (c *Converter) Process(ctx context.Context, localFilePath string, bucket string, object string) error {
	ext, ct := c.ContentType(object)
	if ext == ".md" {
		if err := c.MarkdownToGCS(ctx, localFilePath, bucket, object); err != nil {
			return err
		}
		return nil
	}

	if err := c.LocalToGCS(ctx, localFilePath, bucket, object, ct); err != nil {
		return err
	}

	return nil
}

func (c *Converter) MarkdownToGCS(ctx context.Context, localFilePath string, bucket string, object string) (rerr error) {
	html, err := c.readMarkdownFileToHTML(localFilePath)
	if err != nil {
		return err
	}

	gcsw := c.gcs.NewWriter(ctx, bucket, c.ObjectPath(localFilePath))
	defer func() {
		if err := gcsw.Close(); err != nil {
			rerr = err
		}
	}()
	gcsw.ObjectAttrs.ContentType = "text/html;charset=utf-8"
	_, err = gcsw.Write(html)
	if err != nil {
		return err
	}

	return nil
}

func (c *Converter) readMarkdownFileToHTML(localFilePath string) ([]byte, error) {
	f, err := os.Open(localFilePath)
	if err != nil {
		return nil, err
	}
	md, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	return c.MarkdownToHTML(md), nil
}

func (c *Converter) LocalToGCS(ctx context.Context, localFilePath string, bucket string, object string, contentType string) (rerr error) {
	f, err := os.Open(localFilePath)
	if err != nil {
		return err
	}

	gcsw := c.gcs.NewWriter(ctx, bucket, c.ObjectPath(localFilePath))
	defer func() {
		if err := gcsw.Close(); err != nil {
			rerr = err
		}
	}()

	gcsw.ObjectAttrs.ContentType = contentType
	_, err = io.Copy(gcsw, f)
	if err != nil {
		return err
	}

	return nil
}

// ContentType is ファイル名から拡張子とContentTypeを返す
func (c *Converter) ContentType(fileName string) (ext string, contentType string) {
	ext = filepath.Ext(fileName)

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

	return strings.ToLower(ext), ct
}

func (c *Converter) ObjectPath(localFilePath string) string {
	ret := localFilePath
	if strings.HasPrefix(localFilePath, c.root) {
		ret = localFilePath[len(c.root):]
	}
	if strings.HasPrefix(ret, "/") {
		ret = ret[1:]
	}
	ret = c.extMDToHTML(ret)

	return ret
}

func (c *Converter) extMDToHTML(filePath string) string {
	if strings.HasSuffix(strings.ToLower(filePath), ".md") {
		filePath = filePath[:len(filePath)-3] + ".html"
	}
	return filePath
}
