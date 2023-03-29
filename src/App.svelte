<script>
	import Fa from 'svelte-fa';
	import CryptoJS from 'crypto-js'
	import {faPalette, faGear, faLink} from '@fortawesome/free-solid-svg-icons'
	import EmptyDivs from "./EmptyDivs.svelte";
  import controllers from "./lib/controllers";
	$: haveEvents = 'ongamepadconnected' in window;
	$: gamepads = {}
	$: gamepadConfig = undefined
	$: isDirectionalButton = undefined;
	$: buttonActiveColour = '#5af'
	$: buttonRimColour = 'white';
	$: buttonInactiveColour = undefined;
	$: buttonStyleWindowOpen = false;
	$: inputConfigWindowOpen = false;
	$: inputConfigSetupOngoing = false;
	$: customUrlRetrievalWindowOpen = false;
	$: buttonBorderStyle = `border-color: ${buttonRimColour};`
	const crypto = 'cheaterlikesamir';
	const queryParams = new URLSearchParams(window.location.search);
	const hideConfigParam = queryParams.get('hideConfig');
	const encryptedOptions = queryParams.get('options');
	if (encryptedOptions) {
		const bytes = CryptoJS.AES.decrypt(encryptedOptions, crypto);
		const obj = JSON.parse(bytes.toString(CryptoJS.enc.Utf8));
		gamepadConfig = obj.gamepadConfig;
		buttonActiveColour = obj.buttonActiveColour;
		buttonInactiveColour = obj.buttonInactiveColour;
		buttonRimColour = obj.buttonRimColour;
	}
	const hideConfig = hideConfigParam ? hideConfigParam === 'true' : false;
	const changeButtonInnerStyle = (isButtonTouched) => {
		switch (isButtonTouched) {
			case true:
				return `background-color: ${buttonActiveColour};`
			default:
				if (buttonInactiveColour) {
					return `background-color: ${buttonInactiveColour}`;
				}
				return 'background-color: none';
		}
	}
	const removeGamepad = (index) => delete gamepads[index];
	window.addEventListener("gamepaddisconnected", (e) => {
		removeGamepad(e.gamepad.index);
		gamepadConfig = undefined;
		isDirectionalButton = undefined;
	})
	const addGamepad = (gamepad) => {
		gamepads = {
			...gamepads,
			[gamepad.index]: {
				buttons: gamepad.buttons,
				id: gamepad.id,
				axes: gamepad.axes
			}
		}
		gamepadConfig = controllers.find(controller => {
			return gamepad.id.indexOf(controller.id) !== -1;
		})
	}
	$: if (gamepadConfig && !isDirectionalButton) {
		isDirectionalButton = gamepadConfig.buttons.left.isButton;
	}
	const scanGamepads = () => {
		const detectedGamepads = navigator.getGamepads
      ? navigator.getGamepads()
      : navigator.webkitGetGamepads
      ? navigator.webkitGetGamepads()
      : []
    for (let i = 0; i < detectedGamepads.length; i++) {
      if (detectedGamepads[i]) {
        addGamepad(detectedGamepads[i])
      }
    }
	}

	$: if (!haveEvents) {
		setInterval(scanGamepads, 50);
	}

	window.addEventListener("gamepadconnected", (e) => {
		addGamepad(e.gamepad)
	})
</script>

<main>
	{#if Object.keys(gamepads).length && gamepadConfig}
		{#if !hideConfig}
			<div class="config-buttons"><button on:click={() => {buttonStyleWindowOpen = true;}}><Fa icon={faPalette} /></button> <button on:click={() => {inputConfigWindowOpen = true;}}><Fa icon={faGear} /></button> <button on:click={() => {customUrlRetrievalWindowOpen = true}}><Fa icon={faLink} /></button></div>
		{/if}
		<div class="hitbox-grid">
			<div id="hitbox-up"><div class="button-inner-30" style="{buttonBorderStyle} {changeButtonInnerStyle(isDirectionalButton ? gamepads[0].buttons[gamepadConfig.buttons.up.index].touched : gamepads[0].axes[gamepadConfig.buttons.up.index] === gamepadConfig.buttons.up.value)}" /></div>
			<div id="hitbox-down"><div class="button-inner-24" style="{buttonBorderStyle} {changeButtonInnerStyle(isDirectionalButton ? gamepads[0].buttons[gamepadConfig.buttons.down.index].touched : gamepads[0].axes[gamepadConfig.buttons.down.index] === gamepadConfig.buttons.down.value)}" /></div>
			<div id="hitbox-left"><div class="button-inner-24" style="{buttonBorderStyle} {changeButtonInnerStyle(isDirectionalButton ? gamepads[0].buttons[gamepadConfig.buttons.left.index].touched : gamepads[0].axes[gamepadConfig.buttons.left.index] === gamepadConfig.buttons.left.value)}" /></div>
			<div id="hitbox-right"><div class="button-inner-24" style="{buttonBorderStyle} {changeButtonInnerStyle(isDirectionalButton ? gamepads[0].buttons[gamepadConfig.buttons.right.index].touched : gamepads[0].axes[gamepadConfig.buttons.right.index] === gamepadConfig.buttons.right.value)}" /> </div>
			<div id="hitbox-jab"><div class="button-inner-24" style="{buttonBorderStyle} {changeButtonInnerStyle(gamepads[0].buttons[gamepadConfig.buttons.jab].touched)}" /></div>
			<div id="hitbox-strong"><div class="button-inner-24" style="{buttonBorderStyle} {changeButtonInnerStyle(gamepads[0].buttons[gamepadConfig.buttons.strong].touched)}" /></div>
			<div id="hitbox-fierce"><div class="button-inner-24" style="{buttonBorderStyle} {changeButtonInnerStyle(gamepads[0].buttons[gamepadConfig.buttons.fierce].touched)}" /></div>
			<div id="hitbox-short"><div class="button-inner-24" style="{buttonBorderStyle} {changeButtonInnerStyle(gamepads[0].buttons[gamepadConfig.buttons.short].touched)}" /></div>
			<div id="hitbox-forward"><div class="button-inner-24" style="{buttonBorderStyle} {changeButtonInnerStyle(gamepads[0].buttons[gamepadConfig.buttons.forward].touched)}" /></div>
			<div id="hitbox-roundhouse"><div class="button-inner-24" style="{buttonBorderStyle} {changeButtonInnerStyle(gamepads[0].buttons[gamepadConfig.buttons.roundhouse].touched)}" /></div>
			<div id="hitbox-3p"><div class="button-inner-24" style="{buttonBorderStyle} {changeButtonInnerStyle(gamepads[0].buttons[gamepadConfig.buttons.ppp].touched)}" /></div>
			<div id="hitbox-3k"><div class="button-inner-24" style="{buttonBorderStyle} {changeButtonInnerStyle(gamepads[0].buttons[gamepadConfig.buttons.kkk].touched)}" /></div>
			<EmptyDivs />
		</div>
	{:else}
		<p>Please connect a gamepad.</p>
	{/if}
</main>

<style>
	main {
		text-align: center;
		padding: 1em;
		max-width: 1280px;
		width: 100%;
		max-height: 720px;
		height: 100vh;
		margin: 0 auto;
	}

	.hitbox-grid {
		width: 100%;
		display: grid;
		grid-template-rows: repeat(108, 5px);
		grid-template-columns: repeat(196, 5px);
	}

	.hitbox-grid > div {
		width: 5px;
		height: 5px;
	}	

	#hitbox-left {
		grid-column-start: 1;
		grid-column-end: 24;
		grid-row-start: 12;
		grid-row-end: 35;
	}

	#hitbox-down {
		grid-column-start: 29;
		grid-column-end: 52;
		grid-row-start: 12;
		grid-row-end: 35;
	}

	#hitbox-right {
		grid-column-start: 54;
		grid-column-end: 77;
		grid-row-start: 26;
		grid-row-end: 49;
	}

	#hitbox-up {
		grid-column-start: 69;
		grid-column-end: 98;
		grid-row-start: 79;
		grid-row-end: 105;
	}

	#hitbox-jab {
		grid-column-start: 91;
		grid-column-end: 114;
		grid-row-start: 15;
		grid-row-end: 43;
	}

	#hitbox-short {
		grid-column-start: 91;
		grid-column-end: 114;
		grid-row-start: 44;
		grid-row-end: 73;
	}

	#hitbox-strong {
		grid-column-start: 116;
		grid-column-end: 139;
		grid-row-start: 1;
		grid-row-end: 24;
	}

	#hitbox-forward {
		grid-column-start: 116;
		grid-column-end: 139;
		grid-row-start: 30;
		grid-row-end: 53;
	}

	#hitbox-fierce {
		grid-column-start: 145;
		grid-column-end: 168;
		grid-row-start: 1;
		grid-row-end: 24;
	}

	#hitbox-roundhouse {
		grid-column-start: 145;
		grid-column-end: 168;
		grid-row-start: 30;
		grid-row-end: 53;
	}

	#hitbox-3p {
		grid-column-start: 173;
		grid-column-end: 196;
		grid-row-start: 8;
		grid-row-end: 31;
	}

	#hitbox-3k {
		grid-column-start: 173;
		grid-column-end: 196;
		grid-row-start: 37;
		grid-row-end: 60;
	}

	.button-inner-24 {
		width: 120px;
		height: 120px;
		border-radius: 50%;
	}

	.button-inner-30 {
		width: 150px;
		height: 150px;
		border-radius: 50%;
	}

	button {
		background: none;
	}

	@media (min-width: 640px) {
		main {
			max-width: none;
		}
	}
</style>