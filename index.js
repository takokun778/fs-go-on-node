const { spawn } = require('node:child_process');

const main = spawn('./main');

main.stdout.on('data', (data) => {
    console.log(`stdout: ${data}`);
});

main.stderr.on('data', (data) => {
    console.error(`stderr: ${data}`);
});

main.on('close', (code) => {
    console.log(`child process exited with code ${code}`);
});
