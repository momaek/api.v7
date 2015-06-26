package kodo

import (
	"io"
	"qiniupkg.com/api.v7/kodocli"

	. "golang.org/x/net/context"
)

type PutExtra kodocli.PutExtra

// ----------------------------------------------------------

func (p Bucket) calcUptoken(key string) string {

	policy := &PutPolicy{
		Scope:   p.Name + ":" + key,
		Expires: 3600,
	}
	return policy.Token(p.Conn.mac)
}

func (p Bucket) calcUptokenWithoutKey() string {

	policy := &PutPolicy{
		Scope:   p.Name,
		Expires: 3600,
	}
	return policy.Token(p.Conn.mac)
}

// ----------------------------------------------------------

func (p Bucket) Put(
	ctx Context, ret interface{}, key string, data io.Reader, size int64, extra *PutExtra) error {

	uploader := kodocli.Uploader{Conn: p.Conn.Client, UpHosts: p.Conn.UpHosts}
	uptoken := p.calcUptoken(key)
	return uploader.Put(ctx, ret, uptoken, key, data, size, (*kodocli.PutExtra)(extra))
}

func (p Bucket) PutWithoutKey(
	ctx Context, ret interface{}, data io.Reader, size int64, extra *PutExtra) error {

	uploader := kodocli.Uploader{Conn: p.Conn.Client, UpHosts: p.Conn.UpHosts}
	uptoken := p.calcUptokenWithoutKey()
	return uploader.PutWithoutKey(ctx, ret, uptoken, data, size, (*kodocli.PutExtra)(extra))
}

// ----------------------------------------------------------

func (p Bucket) PutFile(
	ctx Context, ret interface{}, key, localFile string, extra *PutExtra) (err error) {

	uploader := kodocli.Uploader{Conn: p.Conn.Client, UpHosts: p.Conn.UpHosts}
	uptoken := p.calcUptoken(key)
	return uploader.PutFile(ctx, ret, uptoken, key, localFile, (*kodocli.PutExtra)(extra))
}

func (p Bucket) PutFileWithoutKey(
	ctx Context, ret interface{}, localFile string, extra *PutExtra) (err error) {

	uploader := kodocli.Uploader{Conn: p.Conn.Client, UpHosts: p.Conn.UpHosts}
	uptoken := p.calcUptokenWithoutKey()
	return uploader.PutFileWithoutKey(ctx, ret, uptoken, localFile, (*kodocli.PutExtra)(extra))
}

// ----------------------------------------------------------
