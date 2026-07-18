<script setup>
defineProps({
  modelValue: {
    type: Boolean,
    required: true
  },
  newEvent: {
    type: Object,
    required: true
  }
});

defineEmits(['update:modelValue', 'save']);
</script>

<template>
  <v-dialog :model-value="modelValue" @update:model-value="$emit('update:modelValue', $event)" max-width="500px">
    <v-card>
      <v-card-title>Tambah Agenda Baru</v-card-title>
      <v-card-text>
        <v-text-field :model-value="newEvent.name" @update:model-value="newEvent.name = $event" label="Nama Agenda"></v-text-field>
      </v-card-text>
      <v-card-actions>
        <v-spacer></v-spacer>
        <v-btn color="grey" @click="$emit('update:modelValue', false)">Batal</v-btn>
        <v-btn color="primary" @click="$emit('save')">Simpan</v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>
