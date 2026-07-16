import { ref, computed } from 'vue';

export function useEventStatus() {
  const statusDialog = ref(false);
  const itemToChange = ref(null);

  const statusMap = {
    1: { label: 'New', color: 'warning' },
    2: { label: 'Open', color: 'primary' },
    3: { label: 'Completed', color: 'success' },
    4: { label: 'Cancelled', color: 'error' },
  };

  const statusMessage = computed(() => {
    if (!itemToChange.value) return '';
    if (itemToChange.value.status === 1) return 'Promote to Open?';
    if (itemToChange.value.status === 2) return 'Close registration?';
    return 'Change status?';
  });

  const openStatusDialog = (item) => {
    itemToChange.value = item;
    statusDialog.value = true;
  };

  return { statusDialog, itemToChange, statusMessage, statusMap, openStatusDialog };
}
