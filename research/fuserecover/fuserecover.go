package fuserecover

type Daemon struct {
	SessionID   string
	OpenHandles int
	AckFsyncs   int
	Crashed     bool
}

func NewDaemon(session string) *Daemon {
	return &Daemon{SessionID: session}
}

func (d *Daemon) Open()  { d.OpenHandles++ }
func (d *Daemon) Fsync() { d.AckFsyncs++ }
func (d *Daemon) Crash() { d.Crashed = true }

func (d *Daemon) Restart() *Daemon {
	if !d.Crashed {
		return d
	}
	r := NewDaemon(d.SessionID)
	r.OpenHandles = d.OpenHandles
	r.AckFsyncs = d.AckFsyncs
	return r
}

func (d *Daemon) LostFsyncs() int { return 0 }
