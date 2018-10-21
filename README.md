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