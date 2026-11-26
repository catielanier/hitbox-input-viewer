<script>
  import { onMount } from "svelte";
  import { buttons, viewBox } from "./layout.js";

  const query = new URLSearchParams(window.location.search);

  const preview = query.has("preview");
  const showStatus = query.has("status") || preview;

  let connectedToServer = false;
  let controllerConnected = false;
  let pressed = {};
  let connection;

  function togglePreview(input) {
    if (preview) {
      pressed = {
        ...pressed,
        [input]: !pressed[input],
      };
    }
  }

  onMount(() => {
    if (preview) return;

    let disposed = false;
    let reconnectId;

    const connect = () => {
      if (disposed) return;

      const scheme =
        location.protocol === "https:" ? "wss:" : "ws:";

      connection = new WebSocket(
        `${scheme}//${location.host}/ws`
      );

      connection.onopen = () => {
        connectedToServer = true;
      };

      connection.onmessage = ({ data }) => {
        try {
          const packet = JSON.parse(data);

          if (packet.type !== "state") return;

          controllerConnected = Boolean(
            packet.state.connected
          );

          pressed = controllerConnected
            ? packet.state.buttons || {}
            : {};
        } catch (e) {
          console.error("Invalid controller state", e);
        }
      };

      connection.onclose = () => {
        connectedToServer = false;
        controllerConnected = false;
        pressed = {};

        if (!disposed) {
          reconnectId = setTimeout(connect, 1000);
        }
      };

      connection.onerror = () => connection.close();
    };

    connect();

    return () => {
      disposed = true;
      clearTimeout(reconnectId);
      connection?.close();
    };
  });
</script>

<svelte:head>
  <meta name="theme-color" content="#00000000" />
</svelte:head>

<main>
  <svg
    {viewBox}
    role="img"
    aria-label="Leverless controller input display"
    xmlns="http://www.w3.org/2000/svg"
  >
    {#each buttons as button (button.id)}
      <g
        class:active={Boolean(pressed[button.input])}
        class:preview={preview}
        on:click={() => togglePreview(button.input)}
      >
        <circle
          class="rim"
          cx={button.x}
          cy={button.y}
          r={button.r + 4}
        />

        <circle
          class="face"
          cx={button.x}
          cy={button.y}
          r={button.r - 2}
        />

        <text
          x={button.x}
          y={button.y + 1}
          text-anchor="middle"
          dominant-baseline="central"
        >
          {button.label}
        </text>
      </g>
    {/each}
  </svg>

  {#if showStatus}
    <div class="status" class:online={controllerConnected}>
      {#if preview}
        PREVIEW MODE — click buttons
      {:else if controllerConnected}
        CONTROLLER CONNECTED
      {:else if connectedToServer}
        WAITING FOR CONTROLLER
      {:else}
        CONNECTING TO WINDOWS VIEWER
      {/if}
    </div>
  {/if}
</main>

<style>
  :global(html),
  :global(body),
  :global(#app) {
    margin: 0;
    width: 100%;
    height: 100%;
    background: transparent;
    overflow: hidden;
  }

  :global(body) {
    font-family: system-ui, -apple-system, sans-serif;
  }

  main {
    width: 100%;
    height: 100%;
    display: flex;
    align-items: center;
    justify-content: center;
    position: relative;
  }

  svg {
    display: block;
    width: 100%;
    height: 100%;
    overflow: visible;
  }

  g {
    transition: filter 0.045s linear;
  }

  .preview {
    cursor: pointer;
  }

  .rim {
    fill: #0e1725;
    stroke: #b9c5d3;
    stroke-width: 3;
  }

  .face {
    fill: #1b2738;
    stroke: #425169;
    stroke-width: 2;
  }

  text {
    font-weight: 750;
    font-size: 19px;
    fill: #eef5ff;
    pointer-events: none;
    user-select: none;
  }

  g.active .rim {
    fill: #b6f1e6;
    stroke: #f1fffb;
  }

  g.active .face {
    fill: #63d6bf;
    stroke: #d8fff7;
  }

  g.active text {
    fill: #063c34;
  }

  g.active {
    filter: drop-shadow(0 0 8px #6ee7cf);
  }

  .status {
    position: absolute;
    bottom: 8px;
    left: 10px;
    background: #131d2de8;
    color: #f8b4b4;
    border: 1px solid #5d6978;
    border-radius: 6px;
    padding: 5px 8px;
    font-size: 11px;
    font-weight: 700;
    letter-spacing: 0.05em;
  }

  .status.online {
    color: #6ee7b7;
  }
</style>
