package main

import ( 
  "fmt"
  "log"
  "os"
  "time"
  "net" 
  "golang.org/x/crypto/ssh"
)

func connect(user, password, host string, port int) (*ssh.Session, error) { 
  var (
    auth         []ssh.AuthMethod
    addr         string
    clientConfig *ssh.ClientConfig
    client       *ssh.Client
    session      *ssh.Session
    err          error
  )
  // get auth method
  auth = make([]ssh.AuthMethod, 0)
  auth = append(auth, ssh.Password(password))
 
  clientConfig = &ssh.ClientConfig{
    User:    user,
    Auth:    auth,
    Timeout: 30 * time.Second,
    HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
        return nil
    },
  }
 
  // connet to ssh
  addr = fmt.Sprintf("%s:%d", host, port)
 
  if client, err = ssh.Dial("tcp", addr, clientConfig); err != nil {
    return nil, err
  }
 
  // create session
  if session, err = client.NewSession(); err != nil {
    return nil, err
  }
 
  return session, nil
}

func main() { 
  session, err := connect("root", "123456", "10.0.0.66", 22)
  if err != nil {
    log.Fatal(err)
  }
  defer session.Close()
 
  session.Stdout = os.Stdout
  session.Stderr = os.Stderr
  //session.Run("ls /; ls /root")
  session.Run("./a.out")
}
