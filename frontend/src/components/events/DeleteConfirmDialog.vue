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
      <v-card-title class="bg-error text-white">Confirm Delete</v-card-title>
      <v-card-text class="mt-4">
        Are you sure you want to delete "{{ eventName }}"?
      </v-card-text>
      <v-card-actions>
        <v-spacer></v-spacer>
        <v-btn color="grey" @click="$emit('update:modelValue', false)">Cancel</v-btn>
        <v-btn color="error" @click="$emit('confirm')">Yes, Delete</v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>
