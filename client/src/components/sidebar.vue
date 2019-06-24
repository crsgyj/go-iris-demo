<template>
  <div id="sidebar">
    <div class="head"></div>
    <div class="nav">
      <ul class="nav-list">
        <li
          v-for="(item, i) in list"
          :class="['nav-item', {
            active: isActive(item.name)
          }]" 
          :key="i"
          @click="$router.push({ name: item.name })"
        ><span class="nav-link">{{item.text}}</span></li>
      </ul>
    </div>
  </div>
</template>

<script lang="ts">
export default {
  data() {
    return {
      list: [{
        text: '货品检索',
        name:'goods-list'
      }, {
        text: '货品录入',
        name:'goods-write-in'
      }]
    }
  },
  methods: {
    isActive(name) {
      return this.$route.name === name
    }
  }
  
}
</script>

<style lang="stylus" scoped>
#sidebar 
  box-sizing: border-box;
  position: fixed;
  left: 0;
  top: 0;
  bottom: 0;
  width: 220px;
  height: 100vh;
  background #fff
  .head, .nav 
    with 100%
    ul, li 
      list-style none
  .head
    padding 24px 16px

  .nav
    .nav-list
      padding 0
      padding-left 12px
    .nav-item
      display block
      box-sizing border-box
      padding 12px 0
      cursor pointer
      color rgb(32, 45, 64)
      font-size 14px
      text-align center
      border-bottom 2px solid rgba(66, 133, 244, 0)
      // &:hover
      //   background rgb(246, 246, 246)
        
      .nav-link
        display block
      &.active
        background rgb(238, 238, 238)
        color #333
        border-color rgb(66, 133, 244)
        
        
</style>
