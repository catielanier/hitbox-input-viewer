<script>
	let haveEvents = 'ongamepadconnected' in window;
	let gamepads = {}
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
	window.addEventListener("gamepadconnected", (e) => {
		addGamepad(e.gamepad)
	})
</script>

<main>
	{#if Object.keys(gamepads).length}
		<div class="hitbox-grid">
			
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