<template>
  <div>
    <form @submit.prevent="handleSubmit">
      <div class="field">
      <label for="examplePassword" class="label">Enter Password Here:</label>
      <input type="text"
             class="input is-large"
             name="examplePassword"
             id="examplePassword"
             v-model="examplePassword"
             :disabled="showResult.displayResult"
             required>
      </div>
      <button class="button is-success is-large" :disabled="showResult.displayResult">Show me the Hash</button>
    </form>
    <div v-if="showResult.displayResult">
      <div class="message is-danger result-message">
        <div class="message-header">
          <p>The original password was:</p>
        </div>
        <div class="message-body">
          {{examplePassword}}
        </div>
      </div>

      <FinalResult />
      <button @click="resetAll" class="button is-danger is-light is-outlined">Reset</button>
    </div>

  </div>
</template>

<script>
import FinalResult from "@/components/FinalResult";
import {reactive, ref} from 'vue';
import axios from "axios";

export default {
  name: 'Home',
  components: {FinalResult},
  setup() {
    const examplePassword = ref("")
    const showResult = reactive({
      displayResult: false
    })

    const handleSubmit = async () => {
      try {
        await axios.post("http://localhost:3030/api/hashing",{
          examplePassword: examplePassword.value
        })
        showResult.displayResult = true
      } catch (e) {
        console.log(e)
      }

    }

    const resetAll = () => {
      showResult.displayResult = false
      examplePassword.value = ""
    }

    return {
      examplePassword,
      showResult,
      handleSubmit,
      resetAll
    }
  }
}
</script>

<style>
.result-message {
  margin-top: 1.5rem;
}
</style>
