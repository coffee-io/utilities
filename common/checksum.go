package main;

import(
    "crypto/md5"
    "fmt"
    "io"
    "os"
    "path/filepath"
)

func CalculateDirectoryChecksum(directory string) string {
    h := md5.New()
    err := filepath.Walk(directory,
        func(path string, info os.FileInfo, err error) error {
            if info.Mode().IsRegular() {
                f, err := os.Open(path)
                if (err != nil) {
                    return err
                }
                defer f.Close()

                if _, err := io.Copy(h, f); err != nil {
                    return err
                }
            }
            return nil
        })
    if (err != nil) {
        panic(err)
    }
    return fmt.Sprintf("%x", (h.Sum(nil)))
}

// vim:ts=4:sts=4:sw=4:expandtab
