module.exports = {
    apps: [
      {
        name: 'wallet',
        script: 'ts-node',
        args: 'src/server/server.ts',
        interpreter: 'node',
        interpreter_args: '--require ts-node/register',
        watch: false,
        env: {
          NODE_ENV: 'production',
        },
      },
    ],
  };
  