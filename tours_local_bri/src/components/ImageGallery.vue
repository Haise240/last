<template>
  <div class="image-gallery">
    <h2>Image Gallery</h2>
    <div v-if="galleryImages.length">
      <div v-for="(image, index) in galleryImages" :key="index" class="image-item">
        <img :src="image" :alt="'Image ' + (index + 1)" class="gallery-image" />
      </div>
    </div>
    <div v-else>
      <p>No images to display.</p>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      galleryImages: [
        '/static/gallery/319c4ff8b5a6a68fe76d24ff0c32b034.jpg',
        '/static/gallery/345f41244df7b3114240bea6546b914d.jpg',
        '/static/gallery/8eb29048532ec92a068ab07b9d353bac.jpg',
        '/static/gallery/a4164de2f0f82c3949fd9a4f2a58794e.jpg',
      ],
    };
  },
};
</script>

<style scoped>
.image-gallery {
  display: flex;
  flex-wrap: wrap;
}

.image-item {
  margin: 10px;
}

.gallery-image {
  max-width: 200px;
  border: 2px solid #ccc;
  border-radius: 8px;
}
</style>
