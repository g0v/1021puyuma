# 20181021 普悠瑪列車事故

資料來源：[衛生福利部](https://drive.google.com/drive/folders/1qHAOXID3_pW7JBE4xYN2xNjjumKTl38y)

## 手動更新資料

1. 下載 [xlsx](https://drive.google.com/drive/folders/1qHAOXID3_pW7JBE4xYN2xNjjumKTl38y)
2. 更新資料時間 [修改位置](https://github.com/g0v/1021puyuma/blob/gh-pages/parser.js#L47)
3. 執行 `node parse.js`
4. 執行 `git commit` and `git push`

## 也可以使用Golang的Parser
1. `go get github.com/g0v/1021puyuma`
2. 下載 [xlsx](https://drive.google.com/drive/folders/1qHAOXID3_pW7JBE4xYN2xNjjumKTl38y) 
3. 移動xlsx檔案到 `$gopath/src/github.com/g0v/1021puyuma/people.xlsx`
4. 執行 `go run main.go`
5. 執行 `git commit` and `git push`

## 在localhost上測試：
`jekyll serve`

## English: 
On Oct. 21, a Puyuma Express Train [derailed](https://www.theguardian.com/world/2018/oct/21/taiwan-train-derailment-puyuma-express-ludong)
 on the East Coast of Taiwan. The Ministry 
of Health is releasing a series of injured passengers and their whereabouts, but it 
is currently located on a Google Drive link that is periodically updated. 
A group of nobodies at g0v wrote a parser that takes the xlsx file from the Google Drive website
and serves it in a search engine format so the public can easily see who has been affected. 

The source of our data is: 資料來源：[衛生福利部](https://drive.google.com/drive/folders/1qHAOXID3_pW7JBE4xYN2xNjjumKTl38y)

## Updating Data
To update the data on this website, assuming you have Golang installed and your $Gopath set up: 
1. `go get github.com/g0v/1021puyuma`
2. Download this [xlsx](https://drive.google.com/drive/folders/1qHAOXID3_pW7JBE4xYN2xNjjumKTl38y) 
3. Move the xlsx to `$gopath/src/github.com/g0v/1021puyuma/people.xlsx`
4. Run `go run main.go`
5. Run `git commit` and `git push`

## To test on localhost
Run `jekyll serve` on root directory of project.