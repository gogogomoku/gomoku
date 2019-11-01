<template>
	<!-- https://vuejs.org/v2/examples/modal.html -->
	<transition>
    <div class="modal-mask">
      <div class="modal-wrapper">
        <div class="modal-container">

          <div class="modal-header">
            <h3>game settings</h3>
          </div>

          <div class="modal-body">
			  <PlayerSettings
					:id="playerInfo.p1.Id"
				  	:aiStatus="playerInfo.p1.AiStatus"
				/>
				<PlayerSettings
					:id="playerInfo.p2.Id"
				  	:aiStatus="playerInfo.p2.AiStatus"
				/>
				<div class="gameOptions">
					<input type="checkbox"
					id="checkbox"
					:checked="suggestorOn"
					v-bind:disabled="playerInfo.p1.AiStatus > 0 && playerInfo.p2.AiStatus > 0"
					@change="onToggleSuggestor()">
            		<label for="checkbox">enable suggestor</label>
				</div>
        </div>

          <div class="modal-footer">
              <button class="modal-default-button" @click="onClose()">
                OK
              </button>
          </div>
        </div>
      </div>
    </div>
	</transition>
</template>

<script>
import PlayerSettings from './PlayerSettings.vue'

export default {
	name: 'SettingsModal',
	props: {
		"showModal": Boolean,
		"suggestorOn": Boolean,
		"playerInfo": {
			type: Object,
			required: true,
			validator(value) {
				return Object.keys(value).includes("p1", "p2")
			}
		}
	},
	components: {
		PlayerSettings,
	},
	methods: {
		onClose() {
			this.$parent.closeModal()
		},
		onToggleSuggestor() {
			this.$parent.toggleSuggestor()
		}
	}
}

</script>

<style scoped>

.modal-mask {
  position: fixed;
  z-index: 9998;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(0, 0, 0, .5);
  display: table;
  transition: opacity .3s ease;
}

.modal-wrapper {
  display: table-cell;
  vertical-align: middle;
}

.modal-container {
  width: 300px;
  margin: 0px auto;
  padding: 20px 30px;
  background-color:#1c1d21;
  border-radius: 2px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, .33);
  transition: all .3s ease;
}

.modal-header h3 {
  margin-top: 0;
  color: #ece2d0;
}

.modal-body {
  margin: 20px 0;
  color: #bbb;
  font-family:-apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell, 'Open Sans', 'Helvetica Neue', sans-serif;
}

.modal-default-button {
  	float: right;
  	font-family: 'Share Tech Mono', monospace;
}

.modal-footer {
	padding-bottom: 28px;
}

/* button: same as StartButton.vue */
button {
	padding: 9px 12px 9px;
	background-color: #3cadad;
	border: none;
	font-size: 16px;
	font-weight: bolder;
	letter-spacing:normal;
	border-radius: 3px;
}
button:hover {
	background-color: #48cece;
	cursor:pointer;
}
button:active {
	background-color: #2b7c7c;
}

.gameOptions {
	text-align: right;
}

.gameOptions label {
	margin-left: 3px;
}

/*
 * The following styles are auto-applied to elements with
 * transition="modal" when their visibility is toggled
 * by Vue.js.
 *
 * You can easily play with the modal transition by editing
 * these styles.
 */

.modal-enter {
  opacity: 0;
}

.modal-leave-active {
  opacity: 0;
}

.modal-enter .modal-container,
.modal-leave-active .modal-container {
  -webkit-transform: scale(1.1);
  transform: scale(1.1);
}
</style>