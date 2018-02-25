package main

import "testing"

func TestGetFirstCommand(t *testing.T) {
    var commands [4]string
    commands[0] = "ls"
    commands[1] = "ls "
    commands[2] = "ls  -a"
    commands[3] = "ls -la "
    expected := "ls"
    for _, command := range commands {
        result := getFirstComman(command)
        if  result != expected {
            t.Fatalf("Expected %s, but got %s", expected, result)
        }
    }
}

func TestGetTopThreeCommands(t *testing.T) {
    goodCommands := [][]byte{[]byte("git push origin master"),
                             []byte("git push origin master"),
                             []byte("git push origin master"),
                             []byte("git log"),
                             []byte("git status"),
                             []byte("git log"),
                             []byte("git log"),
                             []byte("git status"),
                             []byte("git reflog"),
                             []byte("git log")}
    topThreeCommands := GetTopThreeCommands(goodCommands)
    if len(topThreeCommands) != 3 {
        t.Fatalf("Expected %d commands, but got %d", 3, len(topThreeCommands)
    }
    if topThreeCommands[0] != goodCommands[3] {
        t.Fatalf("Expected %s , but got %s", goodCommands[3], topThreeCommands[0])
    }
    if topThreeCommands[1] != goodCommands[0] {
        t.Fatalf("Expected %s , but got %s", goodCommands[0], topThreeCommands[1])
    }
    if topThreeCommands[2] != goodCommands[4] {
        t.Fatalf("Expected %s , but got %s", goodCommands[4], topThreeCommands[2])
    }
}

func TestGetLessThanThreeCommands(t *testing.T) {
    goodCommands := [][]byte{[]byte("git push origin master"),
                             []byte("git push origin master"),
                             []byte("git push origin master"),
                             []byte("git log")}
    topCommands := GetTopThreeCommands(goodCommands)
    if len(topCommands) != 2 {
        t.Fatalf("Expected %d commands, but got %d", 2, len(topThreeCommands)
    }
    if topCommands[0] != goodCommands[0] {
        t.Fatalf("Expected %s , but got %s", goodCommands[0], topThreeCommands[0])
    }
    if topCommands[1] != goodCommands[3] {
        t.Fatalf("Expected %s , but got %s", goodCommands[3], topThreeCommands[1])
    }
}

func TestNoCommands(t *testing.T) {
    topCommands := GetTopThreeCommands([][]byte{})
    if len(topCommands) != 0 {
        t.Fatalf("Expected %d commands, but got %d", 0, len(topThreeCommands)
    }
}

