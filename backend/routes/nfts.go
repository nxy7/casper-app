package routes

import "net/http"

func NftsByUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("nfts by user"))
}
func NftsAll(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("nfts all"))
}
