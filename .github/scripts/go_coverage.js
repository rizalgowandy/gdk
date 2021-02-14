const util = require("util");
const exec = util.promisify(require("child_process").exec);

async function go_coverage(file) {
  const { stdout } = await exec(
    `go tool cover -func ${file} | grep total | awk '{printf("%s",$3);}'`
  );

  console.log(`${file} Code Coverage:`, stdout);
  return stdout;
}

module.exports = async ({ file }) => {
  return await go_coverage(file);
};
