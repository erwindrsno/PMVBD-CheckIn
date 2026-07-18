<script setup>
defineProps({
  modelValue: {
    type: Boolean,
    required: true
  },
  eventName: {
    type: String,
    required: true
  }
});

defineEmits(['update:modelValue', 'confirm']);
</script>

<template>
  <v-dialog :model-value="modelValue" @update:model-value="$emit('update:modelValue', $event)" max-width="400px">
    <v-card>
      <v-card-title class="bg-error text-white">Konfirmasi Penghapusan</v-card-title>
      <v-card-text class="mt-4">
        Apakah anda yakin ingin menghapus agenda "{{ eventName }}"?
      </v-card-text>
      <v-card-actions>
        <v-spacer></v-spacer>
        <v-btn color="grey" @click="$emit('update:modelValue', false)">Batal</v-btn>
        <v-btn color="error" @click="$emit('confirm')">Ya, Hapus</v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>
