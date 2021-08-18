<template>
  <div class="message is-primary">
    <div class="message-header"><p>The hashed data is: </p></div>
    <div class="message-body">
      {{endResult.resultingHash}}
    </div>
  </div>
</template>

<script>
import {onMounted, reactive} from 'vue'
import axios from "axios";
export default {
  name: "FinalResult",
  setup(){
    const endResult = reactive({
      resultingHash: ""
    })
    onMounted(async () => {
        try {
          const {data} = await axios.get("http://localhost:3030/api/result",)
          endResult.resultingHash = data
        } catch (e) {
          console.log(e)
        }

    })
    return {
      endResult
    }
  }
}
</script>

<style scoped>

</style>