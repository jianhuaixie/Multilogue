package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/jroimartin/gocui"
	peer "github.com/libp2p/go-libp2p-peer"
)

/*
	TODO: Some sort of map of Error code to message

*/

type UIColor struct {
	nameColor string
}

type ChatUI struct {
	node            *Node
	colorMap        map[string]*UIColor
	colorRandomizer *rand.Rand
	channelID       string
	clientPeer      *Peer
	hostPeerID      peer.ID
	isTurn          bool
	host            *Peer
	hostMA          string
}

func NewChatUI(node *Node, channelID string, clientPeer *Peer, hostPeerID peer.ID, hostMA string) *ChatUI {
	var seed = rand.NewSource(time.Now().UnixNano())
	var randomizer = rand.New(seed)

	return &ChatUI{
		node:            node,
		colorMap:        make(map[string]*UIColor),
		colorRandomizer: randomizer,
		channelID:       channelID,
		clientPeer:      clientPeer,
		host:            &Peer{username: "*Channel*"},
		hostPeerID:      hostPeerID,
		hostMA:          hostMA}
}

func (ui *ChatUI) randomColor() string {
	randColor := strconv.Itoa(ui.colorRandomizer.Intn(8))
	bold := ui.colorRandomizer.Intn(2) == 0
	colorString := "\u001b[3" + randColor
	if bold {
		colorString = colorString + ";1"
	}
	return colorString + "m"
}

// Layout creates chat UI
func (ui *ChatUI) Layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	g.Cursor = true

	if messages, err := g.SetView("messages", 0, 0, maxX-1, maxY-5); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		messages.Title = ui.channelID + ": " + ui.hostMA
		messages.Autoscroll = true
		messages.Wrap = true

		messagesView, _ := g.View("messages")
		go ui.BroadcastPanel(g, messagesView, ui.channelID)

	}

	if input, err := g.SetView("input", 0, maxY-5, maxX-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		input.Title = ui.clientPeer.username + ": "
		input.Autoscroll = false
		input.Wrap = true
		input.Editable = true
		if _, err := g.SetCurrentView("input"); err != nil {
			return err
		}
	}

	return nil
}

func (ui *ChatUI) exit(g *gocui.Gui, v *gocui.View) error {
	ui.node.LeaveChannel(ui.clientPeer, ui.hostPeerID, ui.channelID)
	g.Close()
	os.Exit(0)
	return gocui.ErrQuit
}

func (ui *ChatUI) buildMessage(message *Message) string {
	color := ui.randomColor()
	return color + message.peer.username + ": \u001b[0m" + message.text
}

func (ui *ChatUI) sendMessage(g *gocui.Gui, v *gocui.View) error {
	if ui.isTurn {
		msg := strings.TrimSpace(v.Buffer())

		req, _ := ui.node.SendMessage(ui.clientPeer, ui.hostPeerID, ui.channelID, msg)
		var resp *Response
		select {
		case <-req.success:
			// Clear input panel
			g.Update(func(g *gocui.Gui) error {
				v.Clear()
				v.SetCursor(0, 0)
				v.SetOrigin(0, 0)
				return nil
			})
			break
		case resp = <-req.fail:
			ui.isTurn = false
			errorName := ProtocolErrorName(resp.errorCode)
			errorMsg := ui.clientPeer.username + " turn finished because " + errorName
			ui.displayMessage(g, v, ui.buildChannelMessage(errorMsg))
			break
		case <-time.After(3 * time.Second):
			// Timeout
			ui.isTurn = false
			ui.displayMessage(g, v, ui.buildChannelMessage("Message request timed out. "))
			break
		}
	} else {
		ui.displayMessage(g, v, ui.buildChannelMessage("Not your turn. "))
	}

	return nil
}

func (ui *ChatUI) requestTransmission(g *gocui.Gui, v *gocui.View) error {
	if !ui.isTurn {
		req, reqSent := ui.node.SendTransmissionRequest(ui.clientPeer, ui.hostPeerID, ui.channelID)
		if reqSent {
			select {
			case <-req.success:
				ui.isTurn = true
				msg := ui.clientPeer.username + " turn strated  "
				ui.displayMessage(g, v, ui.buildChannelMessage(msg))
			case <-req.fail:
				ui.isTurn = false
				msg := "Can't start turn,  " + ui.clientPeer.username
				ui.displayMessage(g, v, ui.buildChannelMessage(msg))
			}
		} else {
			ui.isTurn = false
			ui.node.debugPrintln("Transmission request not completed.")
		}
	}
	return nil
}

func (ui *ChatUI) endTransmission(g *gocui.Gui, v *gocui.View) error {
	if ui.isTurn {
		ui.node.EndTransmission(ui.clientPeer, ui.hostPeerID, ui.channelID)
		msg := ui.clientPeer.username + " turn ended."
		ui.displayMessage(g, v, ui.buildChannelMessage(msg))
		ui.isTurn = false
	}
	return nil
}

func (ui *ChatUI) StartUI() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Fatal(err)
	}
	defer g.Close()

	g.SetManagerFunc(ui.Layout)
	g.SetKeybinding("input", gocui.KeyEnter, gocui.ModNone, ui.sendMessage)
	g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, ui.exit)
	g.SetKeybinding("", gocui.KeyHome, gocui.ModNone, ui.requestTransmission)
	g.SetKeybinding("", gocui.KeyEnd, gocui.ModNone, ui.endTransmission)
	g.MainLoop()
}

// Displays messages broadcasted to channel
func (ui *ChatUI) BroadcastPanel(g *gocui.Gui, v *gocui.View, channelID string) {
	val, exists := ui.node.channels.Get(channelID)
	if exists == false {
		log.Println("Channel not found.")
		return
	}

	channel := val.(*Channel)

	var message *Message
	for {
		select {
		case message = <-channel.output.messageQueue:
			msgString := "From UI, message reqcieved " + message.text
			ui.node.debugPrintln(msgString)

			ui.displayMessage(g, v, message)
			break
		}
	}
}

func (ui *ChatUI) buildChannelMessage(message string) *Message {
	return &Message{
		peer:      ui.host,
		timestamp: time.Now(),
		text:      message}
}

func (ui *ChatUI) displayMessage(g *gocui.Gui, v *gocui.View, message *Message) {
	color, exists := ui.colorMap[message.peer.username]
	if !exists {
		color = &UIColor{
			nameColor: ui.randomColor()}
		ui.colorMap[message.peer.username] = color
	}

	msg := color.nameColor + message.peer.username +
		"\u001b[0m" + ":" + message.text + "\u001b[0m"

	// Add to message panel
	messagesView, _ := g.View("messages")
	g.Update(func(g *gocui.Gui) error {
		fmt.Fprintln(messagesView, msg)
		return nil
	})
}
