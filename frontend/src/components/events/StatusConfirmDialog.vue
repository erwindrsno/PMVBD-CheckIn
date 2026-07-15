<script setup>
defineProps({
  modelValue: {
    type: Boolean,
    required: true
  },
  statusMessage: {
    type: String,
    required: true
  }
});

defineEmits(['update:modelValue', 'confirm']);
</script>

<template>
  <v-dialog :model-value="modelValue" @update:model-value="$emit('update:modelValue', $event)" max-width="400px">
    <v-card>
      <v-card-title class="bg-warning">Confirm Status Change</v-card-title>
      <v-card-text class="mt-4">
        {{ statusMessage }}
      </v-card-text>
      <v-card-actions>
        <v-spacer></v-spacer>
        <v-btn color="grey" @click="$emit('update:modelValue', false)">Cancel</v-btn>
        <v-btn color="primary" @click="$emit('confirm')">Yes, Confirm</v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>
