class ListNode {
  data: any;
  next: any | null;

  constructor(value: any) {
    this.data = value;
    this.next = null;
  }
}

class LinkedList {
  private head: any | null;
  private tail: any | null;
  private length: number | null;

  constructor() {
    this.head = null;
    this.tail = null;
    this.length = 0;
  }

  //Add an item to the list
  public add(value: any) {
    let newNode = new ListNode(value);

    if (!this.head) {
      this.head = newNode;
      this.tail = newNode;
    } else {
      this.tail.next = newNode;
      this.tail = newNode;
    }
    this.length++;

    return this;
  }

  //Add an item to the beginning of the list
  public addAtStart(value: any) {
    let newNode = new ListNode(value);

    if (this.head === null) {
      this.head = newNode;
      this.tail = newNode;
      this.length++;
      return;
    }

    newNode.next = this.head;
    this.head = newNode;
    this.length++;
  }

  //Add an item to the end of the list
  public addToEnd(value: any) {
    let newNode = new ListNode(value);

    if (this.head === null) {
      this.head = newNode;
      this.tail = newNode;
      this.length++;
      return;
    }

    let current = this.head;

    while (current.next !== null) {
      current = current.next;
    }

    current.next = newNode;
    this.length++;
    this.tail = newNode;
  }

  //Add an item at an index in the list
  public addAtIndex(value: any, index: number) {
    let newNode = new ListNode(value);

    if (index === 0) {
      newNode.next = this.head;
      this.head = newNode;
      this.length++;
    }
    let current = this.head;

    //Start index of the list
    let i = 0;

    //Get the item before the item to be replaced
    while (i < index - 1 && current) {
      current = current.next;
      i++;
    }

    if (current) {
      newNode.next = current.next;
      current.next = newNode;
      this.length++;
    }
  }

  //Deletes the last item
  public delete() {
    let current = this.head;

    if (!current) {
      return null;
    }

    if (this.length === 1) {
      this.head = this.head.next;
      this.length--;
      return this;
    }

    let indexBefore = this.length - 1;

    for (let i = 1; i < indexBefore; i++) {
      current = current.next;
    }

    current.next = current.next.next;

    this.tail = current;

    this.length--;

    return this;
  }

  public deleteAtIndex(index: number) {
    if (index < 0 || index >= this.length) return null;

    if (index === 0) {
      this.head = this.head.next;
      this.length--;
    }

    let current = this.head;

    //Start index of the list
    let i = 0;

    //Get the item before the item to be replaced
    while (i < index - 1 && current) {
      current = current.next;
      i++;
    }

    current.next = current.next.next;

    //If the index of the item to be deleted is the last
    if (index === this.length - 1) {
      this.tail = current;
    }

    this.length--;

    return this;
  }

  //Print Lists of items in the linked list
  public printList() {
    let current = this.head;
    while (current !== null) {
      console.log(current.data);
      current = current.next;
    }
  }

  //Print amount of items in the list
  public getLength() {
    console.log(`The length of the list is: ${this.length}`);
  }

  //Print the start/head and end/tail value of the list
  public getPositions() {
    console.log(
      `The value of the head is: ${this.head.data} and the value of the tail is: ${this.tail.data}`
    );
  }
}

//Test Cases
let linkedList = new LinkedList();

linkedList.add(12);
linkedList.add(2);
linkedList.add(6);
linkedList.add(9);

linkedList.addToEnd(11);

linkedList.addAtStart(8);

linkedList.addAtIndex(1, 2);

linkedList.deleteAtIndex(2);

linkedList.printList();
linkedList.getLength();
linkedList.getPositions();
