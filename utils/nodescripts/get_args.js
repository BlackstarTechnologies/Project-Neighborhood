const args = process.argv.splice(2);
// console.log(args();

const arguments = (() => {
  let t = { args };
  for (let i = 0, end = args.length; i < end; i++) {
    const arg = args[i];
    if (arg.startsWith("-")) {
      t[arg] = i + 1 < end ? args[i + 1] : arg;
    }
  }

  return t;
})();

// console.log(arguments);

module.exports = arguments;
