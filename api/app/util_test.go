package app

import (
	"github.com/globocom/tsuru/fs/testing"
	. "launchpad.net/gocheck"
)

func (s *S) TestFilterOutputWithJujuLog(c *C) {
	output := []byte(`/usr/lib/python2.6/site-packages/juju/providers/ec2/files.py:8: DeprecationWarning: the sha module is deprecated; use the hashlib module instead
  import sha
2012-06-05 17:26:15,881 WARNING ssl-hostname-verification is disabled for this environment
2012-06-05 17:26:15,881 WARNING EC2 API calls not using secure transport
2012-06-05 17:26:15,881 WARNING S3 API calls not using secure transport
2012-06-05 17:26:15,881 WARNING Ubuntu Cloud Image lookups encrypted but not authenticated
2012-06-05 17:26:15,891 INFO Connecting to environment...
2012-06-05 17:26:16,657 INFO Connected to environment.
2012-06-05 17:26:16,860 INFO Connecting to machine 0 at 10.170.0.191
; generated by /sbin/dhclient-script
search novalocal
nameserver 192.168.1.1`)
	expected := []byte(`; generated by /sbin/dhclient-script
search novalocal
nameserver 192.168.1.1`)
	got := filterOutput(output)
	c.Assert(string(got), Equals, string(expected))
}

func (s *S) TestFilterOutputWithoutJujuLog(c *C) {
	output := []byte(`/usr/lib/python2.6/site-packages/juju/providers/ec2/files.py:8: DeprecationWarning: the sha module is deprecated; use the hashlib module instead
  import sha
; generated by /sbin/dhclient-script
search novalocal
nameserver 192.168.1.1`)
	expected := []byte(`; generated by /sbin/dhclient-script
search novalocal
nameserver 192.168.1.1`)
	got := filterOutput(output)
	c.Assert(string(got), Equals, string(expected))
}

func (s *S) TestFiterOutputWithSshWarning(c *C) {
	output := []byte(`2012-06-20 16:54:09,922 WARNING ssl-hostname-verification is disabled for this environment
2012-06-20 16:54:09,922 WARNING EC2 API calls not using secure transport
2012-06-20 16:54:09,922 WARNING S3 API calls not using secure transport
2012-06-20 16:54:09,922 WARNING Ubuntu Cloud Image lookups encrypted but not authenticated
2012-06-20 16:54:09,924 INFO Connecting to environment...
2012-06-20 16:54:10,549 INFO Connected to environment.
2012-06-20 16:54:10,664 INFO Connecting to machine 3 at 10.170.0.166
Warning: Permanently added '10.170.0.121' (ECDSA) to the list of known hosts.
total 0`)
	expected := []byte("total 0")
	got := filterOutput(output)
	c.Assert(string(got), Equals, string(expected))
}

func (s *S) TestFilterOutputWithoutJujuLogAndWarnings(c *C) {
	output := []byte(`; generated by /sbin/dhclient-script
search novalocal
nameserver 192.168.1.1`)
	expected := []byte(`; generated by /sbin/dhclient-script
search novalocal
nameserver 192.168.1.1`)
	got := filterOutput(output)
	c.Assert(string(got), Equals, string(expected))
}

func (s *S) TestFilterOutputRSA(c *C) {
	output := []byte(`/usr/lib/python2.6/site-packages/juju/providers/ec2/files.py:8: DeprecationWarning: the sha module is deprecated; use the hashlib module instead
  import sha
2012-08-22 14:39:18,211 WARNING ssl-hostname-verification is disabled for this environment
2012-08-22 14:39:18,211 WARNING EC2 API calls not using secure transport
2012-08-22 14:39:18,212 WARNING S3 API calls not using secure transport
2012-08-22 14:39:18,212 WARNING Ubuntu Cloud Image lookups encrypted but not authenticated
2012-08-22 14:39:18,222 INFO Connecting to environment...
2012-08-22 14:39:18,854 INFO Connected to environment.
2012-08-22 14:39:18,989 INFO Connecting to machine 4 at 10.170.1.193
Warning: Permanently added '10.170.1.193' (RSA) to the list of known hosts.
Last login: Wed Aug 15 16:08:40 2012 from 10.170.1.239`)
	expected := []byte("Last login: Wed Aug 15 16:08:40 2012 from 10.170.1.239")
	got := filterOutput(output)
	c.Assert(string(got), Equals, string(expected))
}

func (s *S) TestnewUUID(c *C) {
	rfs := &testing.RecordingFs{FileContent: string([]byte{16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31})}
	fsystem = rfs
	defer func() {
		fsystem = s.rfs
	}()
	uuid, err := newUUID()
	c.Assert(err, IsNil)
	expected := "101112131415161718191a1b1c1d1e1f"
	c.Assert(uuid, Equals, expected)
}

func (s *S) TestRandomBytes(c *C) {
	rfs := &testing.RecordingFs{FileContent: string([]byte{16, 17})}
	fsystem = rfs
	defer func() {
		fsystem = s.rfs
	}()
	b, err := randomBytes(2)
	c.Assert(err, IsNil)
	expected := "\x10\x11"
	c.Assert(string(b), Equals, expected)
}
