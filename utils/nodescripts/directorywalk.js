const path = require("path");
const fs = require("fs");

const arguments = require("./get_args");

const { directory, print_t, table, exclude, outFile } = (() => {
  let res = {};
  const proc_cwd = process.cwd();
  const keys = Object.keys(arguments);
  const dir = keys.findIndex((v) => v.match(/dir/));
  const cwd = dir > -1 ? arguments[keys[dir]] : proc_cwd;
  res["directory"] = cwd.endsWith(":") ? cwd + "\\" : cwd;

  const print = keys.findIndex((v) => v.match(/-p/ || /--print_t/));
  res["print_t"] = print > -1 ? true : false;
  res["table"] = { final: [""] };
  const exc = keys.findIndex((v) => v.match(/exc/));
  const cleanedexclude = arguments[keys[exc]].replace('"', "");
  res["exclude"] = exc > -1 ? cleanedexclude : undefined;
  const out_ = keys.findIndex((v) => v.match(/out/));
  const pre_outfile_ = out_ > -1 ? arguments[keys[out_]] : undefined;
  res["outFile"] = pre_outfile_.includes(":")
    ? pre_outfile_
    : path.relative(proc_cwd, pre_outfile_);
  console.log(res);
  return res;
})();

async function Walk(
  options = { dir: ".", table: { final: [""] }, print_t: false }
) {
  const { dir, table, print_t } = options;
  if (typeof dir !== "string") return table;
  const p = path.resolve(__dirname, dir);
  const files_1 = fs
    .readdirSync(p, {
      withFileTypes: true,
    })

    .map((v) => {
      v.name = `${p}\\${v.name}`.replace("\\\\", "\\");
      return v;
    });
  // .forEach((c) => console.log(c));

  const files_2 = files_1;
  console.log(files_2.map((v) => `${v.name}: ${v.isDirectory()}`));

  table.final.push(...files_1.map((v) => v.name));
  const excs = exclude.replace('"', "").split(",");
  // console.log(excs);
  files_2
    .filter((v) => {
      for (let i = 0, end = excs.length; i < end; i++)
        if (v.name.match(excs[i]) || v.name.includes(".")) return false;
      return true;
    })
    .forEach(async (v) => {
      if (v.isDirectory()) {
        await Walk({ dir: v.name, print_t, table, exclude }).catch((e) =>
          console.log("error %d : ", e.errno, e.message)
        );
      }
    });
  // console.log(table);

  return table;
}

Walk({ dir: directory, print_t, table })
  .then((t) => {
    if (outFile)
      fs.writeFileSync(outFile, t.final.join("\n"), {
        encoding: "utf8",
        // flag: "a+",
      });
    console.log(t);
  })
  .catch((e) => console.log("error %d : ", e.errno, e.message));

module.exports = Walk;
