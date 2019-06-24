<template>
  <div class="container" :style="styles"><slot></slot></div>
</template>

<script>
export default {
  data() {
    return {
      showSidebar: true,
    }
  },
  computed: {
    styles() {
      return {
        paddingLeft: this.showSidebar ? '220px' : '0'
      }
    }
  },
  watch: {
    '$route.meta'(val) {  
      this.showSidebar = val && val.sidebar
    }
  },
  
}
</script>

<style lang="stylus" scoped>
.container
  width 100%
  min-height 100vh
  box-sizing border-box
</style>
