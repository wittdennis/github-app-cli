# github-app

github-app is a command line interface tool to interact with github apps.

## Usage

Create a config file `$HOME/.github-app/config.yaml` with the following content:

```yaml
appId: <id of the github app>
privateKey: <private key of the app>
```

After that you can create a access token by invoking:

```sh
githup-app generate-token
```

You can also generate an installation token by invoking:

```sh
github-app generate-token installation INSTALLATION_ID
```
