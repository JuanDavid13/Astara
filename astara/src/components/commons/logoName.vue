<template>
  <div id="logoName">
    <svg :width="logoSize" :height="logoSize" version="1.1" viewBox="0 0 46.029 46" xmlns="http://www.w3.org/2000/svg" xmlns:cc="http://creativecommons.org/ns#" xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:rdf="http://www.w3.org/1999/02/22-rdf-syntax-ns#">
 <!--<metadata>
  <rdf:RDF>
   <cc:Work rdf:about="">
    <dc:format>image/svg+xml</dc:format>
    <dc:type rdf:resource="http://purl.org/dc/dcmitype/StillImage"/>
    <dc:title/>
   </cc:Work>
  </rdf:RDF>
 </metadata>-->
 <g id="logoG" transform="translate(-67.199 -88.399)" fill-opacity="0" stroke-width="2">
  <circle transform="rotate(-90)" cx="-111.4" cy="90.229" r="22"/>
  <ellipse cx="90.229" cy="111.4" rx="9.7289" ry="9.7289"/>
  <path transform="rotate(-90)" d="m-104.93 82.991a22 22 0 0 1-6.4733-15.097"/>
  <path transform="rotate(-90)" d="m-111.41 112.22a22 22 0 0 1 6.1568-14.466"/>
  <path transform="rotate(-90)" d="m-111.37 67.199a22 22 0 0 1-6.5907 15.702 22 22 0 0 1-15.823 6.2941"/>
  <path transform="rotate(-90)" d="m-133.97 91.135a22 22 0 0 1 22.561 20.993"/>
 </g>
</svg>
  <span :style="fontSize">Astara</span>
  </div>
</template>

<script>
export default {
  name:'logoName',
  props: ['size'],
  computed: {
    logoSize(){ return this.size + 1 +"rem"; },
    fontSize(){ return "font-size: " + this.size + "rem;"; },
  }
}
</script>

<style scoped lang="scss">
@import '@/assets/style/common';

#logoName{
  display:flex;
  flex-direction:row;
  align-items:center;
  flex-wrap:nowrap;
  color:var(--gold);
  gap:1rem;
  font-weight:bolder;
}

#logoG{
  stroke:var(--gold);
}
</style>
