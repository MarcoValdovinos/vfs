package fuserecover

import "testing"

func TestH1_OpenHandlesRecovered(t *testing.T) {
	d := NewDaemon("sess-1")
	d.Open()
	d.Open()
	d.Open()
	d.Crash()
	r := d.Restart()
	if r.OpenHandles != 3 {
		t.Fatalf("H1 fallo: esperaba 3 handles, got %d", r.OpenHandles)
	}
	if r.SessionID != "sess-1" {
		t.Fatalf("H1 fallo: session no preservada")
	}
}

func TestH2_ZeroLossFsync(t *testing.T) {
	d := NewDaemon("sess-1")
	for i := 0; i < 10; i++ {
		d.Fsync()
	}
	d.Crash()
	r := d.Restart()
	if r.AckFsyncs != 10 {
		t.Fatalf("H2 fallo: esperaba 10 fsyncs, got %d", r.AckFsyncs)
	}
	if r.LostFsyncs() != 0 {
		t.Fatalf("H2 fallo: perdida debe ser 0")
	}
	t.Logf("H2 OK: ack=10 lost=0")
}

func TestH3_NoCrashNoRestart(t *testing.T) {
	d := NewDaemon("sess-1")
	d.Open()
	d.Fsync()
	r := d.Restart()
	if r != d {
		t.Fatalf("H3 fallo: sin crash no debe crear nuevo daemon")
	}
}

func TestEdge_EmptySession(t *testing.T) {
	d := NewDaemon("sess-1")
	d.Crash()
	r := d.Restart()
	if r.OpenHandles != 0 {
		t.Fatalf("borde fallo: esperaba 0")
	}
}
