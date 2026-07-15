<script setup>
defineProps({
  modelValue: { type: Boolean, required: true },
  selectedAttendee: { type: Object, default: null }
});

// Added 'print-qr' to the emits list
defineEmits(['update:modelValue', 'print-qr']);
</script>

<template>
  <v-dialog :model-value="modelValue" @update:model-value="$emit('update:modelValue', $event)" max-width="600px">
    <v-card v-if="selectedAttendee">
      <v-card-title class="bg-primary text-white">Attendee Details</v-card-title>
      <v-card-text class="mt-4">
        <v-list density="compact">
          <v-list-item title="Name" :subtitle="selectedAttendee.name"></v-list-item>
          <v-list-item title="ID" :subtitle="selectedAttendee.public_id"></v-list-item>
          <v-list-item title="School" :subtitle="selectedAttendee.school"></v-list-item>
          <v-list-item title="Grade" :subtitle="`${selectedAttendee.grade} - ${selectedAttendee.subgrade}`"></v-list-item>
          <v-list-item title="Phone" :subtitle="selectedAttendee.telp"></v-list-item>
          <v-list-item title="Guardian Phone" :subtitle="selectedAttendee.guardian_telp"></v-list-item>
        </v-list>
      </v-card-text>

      <v-card-actions>
        <!-- New Print QR Button -->
        <v-btn
          variant="flat"
          color="secondary"
          prepend-icon="mdi-qrcode"
          @click="$emit('print-qr', selectedAttendee)"
        >
          Print QR
        </v-btn>

        <v-spacer></v-spacer>
        <v-btn color="primary" @click="$emit('update:modelValue', false)">Close</v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>
