package main;

import (
    "testing"
)

func TestChecksumEq(t *testing.T) {
    var a, b = CalculateDirectoryChecksum("."), CalculateDirectoryChecksum(".");
    t.Log("Checksum .:", b);
    if a != b {
        t.Errorf("Different checksums for same directory.");
    }
}

func TestChecksumDif(t *testing.T) {
    var a, b = CalculateDirectoryChecksum("."), CalculateDirectoryChecksum("..");
    t.Log("Checksum ..:", b);
    if a == b {
        t.Errorf("Same checksum for different directories.");
    }
}

// vim:ts=4:sts=4:sw=4:expandtab
