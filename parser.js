const fs = require('fs');
const _ = require('lodash');
const XLSX = require("xlsx");

const xlsx = XLSX.readFile('people.xlsx');
const sheetName = xlsx.SheetNames[0];
const sheet = xlsx.Sheets[sheetName];
const data = [];

const name = {
  A: '編號',
  B: '縣市別',
  C: '收治單位',
  D: '檢傷編號',
  E: '姓名',
  F: '性別',
  G: '國籍',
  H: '年齡',
  I: '醫療檢傷',
  J: '救護檢傷',
  K: '即時動向',
  L: '轉診要求',
  M: '刪除註記',
};

_.forEach(sheet, (row, key) => {
  const idx = /([A-Z])(\d+)/i.exec(key);
  if (!idx) return;
  if (idx[2] <= 2) return;

  const i = parseInt(idx[2], 10);
  if (!data[i]) data[i] = {};
  data[i][name[idx[1]]] = _.trim(row.v);
});

const result = _.map(
  _.filter(data, _.identity),
  (row) => {
    _.forEach(name, (field) => {
      if (!row[field]) row[field] = '';
    });
    return row;
  },
);

fs.writeFileSync('people.json', JSON.stringify({
  lastmodify: "2018/10/21 21:55",
  license: "CC0",
  source: "衛生福利部 https://drive.google.com/drive/folders/1qHAOXID3_pW7JBE4xYN2xNjjumKTl38y",
  data: result,
}, null, 2));
