package main

import (
    "encoding/json"
    "os"
    "sync"
    "time"
)

type GuestbookEntry struct {
    Name      string `json:"name"`
    Timestamp string `json:"timestamp"`
}

var (
    guestbook     []GuestbookEntry
    guestbookLock sync.Mutex
    guestbookFile = "guestbook.json"
)

// Loads entries from disk on startup
func loadGuestbook() {
    data, err := os.ReadFile(guestbookFile)
    if err != nil {
        guestbook = []GuestbookEntry{}
        return
    }
    json.Unmarshal(data, &guestbook)
}

// Saves entries to disk
func saveGuestbook() {
    data, _ := json.MarshalIndent(guestbook, "", "  ")
    os.WriteFile(guestbookFile, data, 0644)
}

// Adds a new entry
func addGuestbookEntry(name string) GuestbookEntry {
    guestbookLock.Lock()
    defer guestbookLock.Unlock()

    entry := GuestbookEntry{
        Name:      name,
        Timestamp: time.Now().Format("2006-01-02 15:04:05"),
    }
    guestbook = append(guestbook, entry)
    saveGuestbook()
    return entry
}

// Returns all entries
func getGuestbookEntries() []GuestbookEntry {
    guestbookLock.Lock()
    defer guestbookLock.Unlock()
    return guestbook
}
