# shipator

Tool to ship static files, injects env vars at runtime and start a nginx server

## Motivation

The common react build tools inject the env vars at build time. Because of that we need to build multiple times for different environments.
This project aims to feel that gap by using the [approach mentioned at the `create-react-app` documentation](https://create-react-app.dev/docs/title-and-meta-tags#injecting-data-from-the-server-into-the-page).

## How it works

The shipator script runs right after container starts, it fetches the env vars that start with `REACT_APP` and the `NODE_ENV` env var and injects it in the target file, replacing the placeholder with the runtime env vars. Then it starts a nginx server serving the static files.

## Usage

1. Update the `public/index.html` of your project, inserting the `__ENV__` to be injected at runtime, but fallbacking to the build time env vars that create-react-app provides, so it work locally normally.

```html
<script>
  window.ENV = (function() {
    try {
      return __ENV__;
    } catch (e) {
      return {};
    }
  })();
  window.ENV.REACT_APP_SERVER_URL =
    window.ENV.REACT_APP_SERVER_URL || '%REACT_APP_SERVER_URL%';
</script>
```

2. Update the rest of the project to get the values from this global variable:

```ts
const baseUrl: string = (window as any).ENV.REACT_APP_SERVER_URL!;
```

3. Now in the k8s configmap we need to set the shipator env variables.

```
SHIPATOR_TARGET: '/usr/share/nginx/html/index.html'
SHIPATOR_PLACEHOLDER: '__ENV__'
```

4. Now all the env vars start with `REACT_APP` or the `NODE_ENV` itself will replace the `__ENV__` when the container starts.
