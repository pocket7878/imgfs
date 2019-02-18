package session

import "sync"
import "context"
import "github.com/docker/go-p9p"

type session struct {
	sync.Mutex
}

func NewSession(ctx context.Context) p9p.Session {
	return &session{}
}

func (sess *session) Auth(ctx context.Context, afid p9p.Fid, uname, aname string) (p9p.Qid, error) {
	return p9p.Qid{}, nil
}

func (sess *session) Attach(ctx context.Context, fid, afid p9p.Fid, uname, aname string) (p9p.Qid, error) {
	return p9p.Qid{}, nil
}

func (sess *session) Clunk(ctx context.Context, fid p9p.Fid) error {
	return nil
}

func (sess *session) Create(ctx context.Context, parent p9p.Fid, name string, perm uint32, mode p9p.Flag) (p9p.Qid, uint32, error) {
	return p9p.Qid{}, 0, nil
}

func (sess *session) Open(ctx context.Context, fid p9p.Fid, mode p9p.Flag) (p9p.Qid, uint32, error) {
	return p9p.Qid{}, 0, nil
}

func (sess *session) Read(ctx context.Context, fid p9p.Fid, p []byte, offset int64) (n int, err error) {
	return 0, nil
}

func (sess *session) Write(ctx context.Context, fid p9p.Fid, p []byte, offset int64) (n int, err error) {
	return 0, nil
}

func (sess *session) Remove(ctx context.Context, fid p9p.Fid) error {
	return nil
}

func (sess *session) Stat(ctx context.Context, fid p9p.Fid) (p9p.Dir, error) {
	return p9p.Dir{}, nil
}
func (sess *session) WStat(ctx context.Context, fid p9p.Fid, dir p9p.Dir) error {
	return nil
}

func (sess *session) Version() (msize int, version string) {
	return 0, ""
}

func (sess *session) Walk(ctx context.Context, fid p9p.Fid, newfid p9p.Fid, names ...string) ([]p9p.Qid, error) {
	qs := make([]p9p.Qid, 0)
	return qs, nil
}
