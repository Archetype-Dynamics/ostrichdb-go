package lib

import "net/http"


//Wrote this since Golang doesnt have one for some fucking reason...
func Delete(path string) (*http.Response, error) {
    client := &http.Client{}
    request, err := http.NewRequest("DELETE", path, nil)
    if err != nil {
        return nil, err
    }
    return client.Do(request)
}

func Put(path string) (*http.Response, error) {
	client := &http.Client{}
    request, err := http.NewRequest("PUT", path, nil)
    if err != nil {
        return nil, err
    }
    return client.Do(request)
}