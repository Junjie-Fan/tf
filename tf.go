package tf

import (
	"sync"

	"github.com/Junjie-Fan/tfidf"
)

type Callpus struct {
	tf    *tfidf.TFIDF
	rwtex sync.RWMutex
}

func New() *Callpus {
	return &Callpus{
		tf: tfidf.New(),
	}
}

func (callpus *Callpus) AddCorpus(s string) {
	callpus.rwtex.Lock()
	defer callpus.rwtex.Unlock()
	callpus.tf.AddDocs(s)
}

func (callpus *Callpus) TargetInit(i int) {
	callpus.tf.InitTerms(i)
	return
}
