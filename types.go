package GoIPerf

var Location = "iperf3"

// Result contains the result of the IPerf3 test
type Result struct {
	Initial  *Initial     `json:"start"`
	Measures *[]Measure   `json:"intervals"`
	Result   *FinalResult `json:"end"`
	Error    string       `json:"error"`
}

// IsBusy is a shortcut for error handling, it returns true if the server is busy
func (r *Result) IsBusy() bool {
	return r.Error == "error - the server is busy running a test. try again later"
}

// IsTemporarilyDown is a shortcut for error handling, it returns true if the server didn't reply to the control message
// Intended for short outages that could be fixed in a few seconds of with another try
func (r *Result) IsTemporarilyDown() bool {
	return r.Error == "error - unable to receive control message: Connection reset by peer"
}

// IsDown is a shortcut for error handling, it returns true if the server is down or unreachable
// Intended for long outages that can't autorecover in a small timespan
func (r *Result) IsDown() bool {
	return r.Error == "error - unable to connect to server: Connection timed out"
}

// Initial contains data about the IPerf3 initialization
type Initial struct {
	Connected             []map[string]interface{} `json:"connected"`
	Version               string                   `json:"version"`
	SystemInfo            string                   `json:"system_info"`
	TimeStamp             *TimeStamp               `json:"timestamp"`
	Server                *Server                  `json:"connecting_to"`
	Cookie                string                   `json:"cookie"`
	TCPMSegmentSize       int                      `json:"tcp_mss_default"`
	SocketBufferSize      int                      `json:"sock_bufsize"`
	SenderBufferCurrent   int                      `json:"sndbuf_actual"`
	ReceiverBufferCurrent int                      `json:"rcvbuf_actual"`
	Test                  *Test                    `json:"test_start"`
}

// TimeStamp contains the timestamp of the test in different formats
type TimeStamp struct {
	TimeStamp string `json:"time"`
	Time      int    `json:"timesecs"`
}

// Server contains data about the destination server
type Server struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

// Test contains data about the IPerf3 test
type Test struct {
	Protocol      string `json:"protocol"`
	StreamsNumber int    `json:"num_streams"`
	BlockSize     int    `json:"blksize"`
	Omit          int    `json:"omit"`
	Duration      int    `json:"duration"`
	Bytes         int    `json:"bytes"`
	Blocks        int    `json:"blocks"`
	Reverse       int    `json:"reverse"`
	ToS           int    `json:"tos"`
}

// Measure contains data about the test Measurements
type Measure struct {
	Streams *[]Stream `json:"streams"`
	Summary *Summary  `json:"sum"`
}

// Stream contains data about the streams used in the test
type Stream struct {
	Socket        int     `json:"socket"`
	Start         float64 `json:"start"`
	End           float64 `json:"end"`
	Seconds       float64 `json:"seconds"`
	BitsPerSecond float64 `json:"bits_per_second"`
	Retransmits   int     `json:"retransmits"`
	SenderCWND    int     `json:"snd_cwnd"`
	RTT           int     `json:"rtt"`
	RTTVar        int     `json:"rttvar"`
	PMTU          int     `json:"pmtu"`
	Omitted       bool    `json:"omitted"`
}

// Summary contains a recap of the whole test
type Summary struct {
	Start         float64 `json:"start"`
	End           float64 `json:"end"`
	Seconds       float64 `json:"seconds"`
	Bytes         int     `json:"bytes"`
	BitsPerSecond float64 `json:"bits_per_second"`
	Retransmits   int     `json:"retransmits"`
	Omitted       bool    `json:"omitted"`
}

// GetMegabits returns the result speed in Megabits instead of bits
func (s *Summary) GetMegabits() float64 {
	return s.BitsPerSecond / 1000000
}

// FinalResult contains all the necessary infos of the test
type FinalResult struct {
	Streams               *[]Measure `json:"streams"`
	Sent                  *Summary   `json:"sum_sent"`
	Received              *Summary   `json:"sum_received"`
	CPU                   *CPU       `json:"cpu_utilization_percent"`
	SenderTCPCongestion   string     `json:"sender_tcp_congestion"`
	ReceiverTCPCongestion string     `json:"receiver_tcp_congestion"`
}

// CPU contains data about the server/client CPU usage
type CPU struct {
	HostTotal    float64 `json:"host_total"`
	HostUser     float64 `json:"host_user"`
	HostSystem   float64 `json:"host_system"`
	RemoteTotal  float64 `json:"remote_total"`
	RemoteUser   float64 `json:"remote_user"`
	RemoteSystem float64 `json:"remote_system"`
}

// IsOverLoaded returns true if the server or the host had too high CPU usage during the tests
func (c *CPU) IsOverLoaded() bool {
	return c.HostTotal > 80 || c.RemoteTotal > 80
}
