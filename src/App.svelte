<script>
	import EmptyDivs from "./EmptyDivs.svelte";
	$: haveEvents = 'ongamepadconnected' in window;
	$: gamepads = {}
	const removeGamepad = (index) => delete gamepads[index];
	window.addEventListener("gamepaddisconnected", (e) => {
		removeGamepad(e.gamepad.index);
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
	}
	const scanGamepads = () => {
		var detectedGamepads = navigator.getGamepads
      ? navigator.getGamepads()
      : navigator.webkitGetGamepads
      ? navigator.webkitGetGamepads()
      : []

    // Loop through all detected controllers and add if not already in state
    for (var i = 0; i < detectedGamepads.length; i++) {
      if (detectedGamepads[i]) {
        addGamepad(detectedGamepads[i])
      }
    }
	}

	$: if (!haveEvents) {
		setInterval(scanGamepads, 50)
	}

	window.addEventListener("gamepadconnected", (e) => {
		addGamepad(e.gamepad)
	})
</script>

<main>
	{#if Object.keys(gamepads).length}
		<div class="hitbox-grid">
			<div id="hitbox-up"><div class="button-inner-30 button-active-{gamepads[0].buttons[12].touched}" /></div>
			<div id="hitbox-down"><div class="button-inner-24 button-active-{gamepads[0].buttons[13].touched}" /></div>
			<div id="hitbox-left"><div class="button-inner-24 button-active-{gamepads[0].buttons[14].touched}" /></div>
			<div id="hitbox-right"><div class="button-inner-24 button-active-{gamepads[0].buttons[15].touched}" /> </div>
			<div id="hitbox-jab"><div class="button-inner-24 button-active-{gamepads[0].buttons[2].touched}" /></div>
			<div id="hitbox-strong"><div class="button-inner-24 button-active-{gamepads[0].buttons[3].touched}" /></div>
			<div id="hitbox-fierce"><div class="button-inner-24 button-active-{gamepads[0].buttons[5].touched}" /></div>
			<div id="hitbox-short"><div class="button-inner-24 button-active-{gamepads[0].buttons[0].touched}" /></div>
			<div id="hitbox-forward"><div class="button-inner-24 button-active-{gamepads[0].buttons[1].touched}" /></div>
			<div id="hitbox-roundhouse"><div class="button-inner-24 button-active-{gamepads[0].buttons[7].touched}" /></div>
			<div id="hitbox-3p"><div class="button-inner-24 button-active-false" /></div>
			<div id="hitbox-3k"><div class="button-inner-24 button-active-false" /></div>
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
		border: 5px solid black;
		border-radius: 50%;
	}

	.button-inner-30 {
		width: 150px;
		height: 150px;
		border: 5px solid black;
		border-radius: 50%;
	}

	.button-active-true {
		background-color: #5af;
	}


	@media (min-width: 640px) {
		main {
			max-width: none;
		}
	}
</style>