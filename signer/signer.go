package main

import (
    "fmt"
    "sort"
    "strconv"
    "sync"
)

const BufferSize = MaxInputDataLen
const MultiHashLength = 6

var md5Mutex sync.Mutex

func ToString(dataRaw interface{}) string {
    strVal, ok := dataRaw.(string)
    if ok {
        return strVal
    }

    intVal, ok := dataRaw.(int)
    if ok {
       return strconv.FormatInt(int64(intVal), 10)
    }

    panic(fmt.Sprintf("impossible convert %T to string", dataRaw))
}

func CountMd5(data string) string {
    md5Mutex.Lock()
    defer md5Mutex.Unlock()
    return DataSignerMd5(data)
}

func CountCrc32(data string, out chan<- string) {
    out <- DataSignerCrc32(data)
}

func SingleHash(in, out chan interface{}) {
    wg := &sync.WaitGroup{}
    for v := range in {
        data := ToString(v)

        results := make([]chan string, 2)
        for i := range results {
            results[i] = make(chan string)
        }

        // In fact, it’s not necessary to use thread-safe 'CountMd5', because calls are strictly sequential.
        go CountCrc32(data, results[0])
        go CountCrc32(CountMd5(data), results[1])

        wg.Add(1)
        go func() {
            defer wg.Done()
            out <- (<-results[0]) + "~" + (<-results[1])
        }()
    }
    wg.Wait()
}

func MultiHash(in, out chan interface{}) {
    wg := &sync.WaitGroup{}
    for v := range in {
        data := ToString(v)

        results := make([]chan string, MultiHashLength)
        for i := range results {
            results[i] = make(chan string)
            go CountCrc32(strconv.FormatInt(int64(i), 10) + data, results[i])
        }

        wg.Add(1)
        go func() {
            defer wg.Done()
            var result string
            for i := range results {
                result += <-results[i]
            }
            out <- result
        }()
    }
    wg.Wait()
}

func CombineResults(in, out chan interface{}) {
    var strings []string
    for v := range in {
        strings = append(strings, ToString(v))
    }
    sort.Strings(strings)

    var result string
    for _, str := range strings {
        result += str + "_"
    }
    // Delete '_' in the end
    result = result[:len(result)-1]
    out <- result
}

func ExecuteFunc(job job, in, out chan interface{}, wg *sync.WaitGroup) {
    defer wg.Done()
    job(in, out)
    close(out)
}

func ExecutePipeline(jobs ...job) {
    wg := &sync.WaitGroup{}
    in := make(chan interface{}, BufferSize)
    close(in)

    for _, job := range jobs {
        out := make(chan interface{}, BufferSize)
        wg.Add(1)
        go ExecuteFunc(job, in, out, wg)
        in = out
    }

    wg.Wait()
}
