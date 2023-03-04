<script>
	import EmptyDivs from "./EmptyDivs.svelte";
	$: haveEvents = 'ongamepadconnected' in window;
	$: gamepads = {}
	const removeGamepad = (index) => delete gamepads.index;
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
		console.log(gamepads)
	}

	$: if (!haveEvents) {
		setInterval(scanGamepads, 10000)
	}
	window.addEventListener("gamepadconnected", (e) => {
		addGamepad(e.gamepad)
	})
</script>

<main>
	{#if Object.keys(gamepads).length}
		<div class="hitbox-grid">
			<div id="hitbox-up"><div class="button-inner button-active-{gamepads[0].buttons[12].touched}" /></div>
			<div id="hitbox-down"><div class="button-inner button-active-{gamepads[0].buttons[13].touched}" /></div>
			<div id="hitbox-left"><div class="button-inner button-active-{gamepads[0].buttons[14].touched}" /></div>
			<div id="hitbox-right"><div class="button-inner button-active-{gamepads[0].buttons[15].touched}" /> </div>
			<div id="hitbox-jab"><div class="button-inner button-active-{gamepads[0].buttons[2].touched}" /></div>
			<div id="hitbox-strong"><div class="button-inner button-active-{gamepads[0].buttons[3].touched}" /></div>
			<div id="hitbox-fierce"><div class="button-inner button-active-{gamepads[0].buttons[5].touched}" /></div>
			<div id="hitbox-short"><div class="button-inner button-active-{gamepads[0].buttons[0].touched}" /></div>
			<div id="hitbox-forward"><div class="button-inner button-active-{gamepads[0].buttons[1].touched}" /></div>
			<div id="hitbox-roundhouse"><div class="button-inner button-active-{gamepads[0].buttons[7].touched}" /></div>
			<div id="hitbox-3p"><div class="button-inner button-active-false" /></div>
			<div id="hitbox-3k"><div class="button-inner button-active-false" /></div>
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
		max-width: 240px;
		margin: 0 auto;
	}

	.hitbox-grid {
		width: 100%;
		display: grid;
		grid-template-rows: repeat(105, 1fr);
		grid-template-columns: repeat(172, 1fr);
	}


	@media (min-width: 640px) {
		main {
			max-width: none;
		}
	}
</style>