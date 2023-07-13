<template>
  <div class="row">
    <div class="span24-8 px-1 mb-1" v-for="(item,index) in label.values" :key="index">
      <div class="rounded bg-light-secondary font text-center border" style="padding:5px;"
           :class="label.selected === item ? 'radio-active':'border-light-secondary'" @click="changeRadio(item)">
        {{item}}
      </div>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    label: Object,
  },
  methods: {
    changeRadio(index) {
      this.$emit('update:selected', index)
      this.$emit('change', index)
    }
  },
}
</script>

<style scoped>
.radio-active {
  background: #fce0d5 !important;
  color: #eb7320 !important;
  border-color: #eb7320 !important;
}
</style>
