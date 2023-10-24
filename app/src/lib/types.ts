export type Page = {
	url: string;
	breadcrumb: string;
	title: string;
	content: string;
};

export type SearchResult = {
	page: Page;
	highlight: string;
};

type EventHandler<E extends Event = Event, T extends EventTarget = Element> =
	// eslint-disable-next-line @typescript-eslint/no-explicit-any
	(event: E & { currentTarget: EventTarget & T }) => any;

export type ClipboardEventHandler<T extends EventTarget> = EventHandler<ClipboardEvent, T>;
export type CompositionEventHandler<T extends EventTarget> = EventHandler<CompositionEvent, T>;
export type DragEventHandler<T extends EventTarget> = EventHandler<DragEvent, T>;
export type FocusEventHandler<T extends EventTarget> = EventHandler<FocusEvent, T>;
export type FormEventHandler<T extends EventTarget> = EventHandler<Event, T>;
export type ChangeEventHandler<T extends EventTarget> = EventHandler<Event, T>;
export type KeyboardEventHandler<T extends EventTarget> = EventHandler<KeyboardEvent, T>;
export type MouseEventHandler<T extends EventTarget> = EventHandler<MouseEvent, T>;
export type TouchEventHandler<T extends EventTarget> = EventHandler<TouchEvent, T>;
export type PointerEventHandler<T extends EventTarget> = EventHandler<PointerEvent, T>;
export type UIEventHandler<T extends EventTarget> = EventHandler<UIEvent, T>;
export type WheelEventHandler<T extends EventTarget> = EventHandler<WheelEvent, T>;
export type AnimationEventHandler<T extends EventTarget> = EventHandler<AnimationEvent, T>;
export type TransitionEventHandler<T extends EventTarget> = EventHandler<TransitionEvent, T>;
export type MessageEventHandler<T extends EventTarget> = EventHandler<MessageEvent, T>;
