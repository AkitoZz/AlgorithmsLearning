#include <stdlib.h>

#define STACKSIZE 100

// 静态数组实现的stack，top值指向栈顶的下标，默认应该是-1
typedef struct {
    int Data[STACKSIZE];
    int Top;
}stack ;

// 两个栈实现的队列
typedef struct {
    stack input;
    stack output;
} MyQueue;

//创建操作，需要初始化stack
MyQueue* myQueueCreate() {
    MyQueue *tmp = (MyQueue*)malloc(sizeof(MyQueue));
    if (NULL == tmp)
        return NULL;
    tmp->input.Top = -1;
    tmp->output.Top = -1;
    return tmp;
}

// 没有做多余的非法判断
void myQueuePush(MyQueue* obj, int x) {
    if (obj->input.Top < STACKSIZE){
        int tmp = obj->input.Top + 1;
        obj->input.Data[tmp] = x;
        obj->input.Top = tmp;
    }
}

// pop时看output栈是否为空
int myQueuePop(MyQueue* obj) {
    if (obj->output.Top != -1){
        int tmp = obj->output.Data[obj->output.Top--];
        return tmp;
    } else {
        while (obj->input.Top != -1){
            //注意此处的先加后复制，和先复制后减
            obj->output.Data[++obj->output.Top] = obj->input.Data[obj->input.Top--];
        }

        return obj->output.Data[obj->output.Top--];
    }

}

// 逻辑同pop，只是不移动下标
int myQueuePeek(MyQueue* obj) {
    if (obj->output.Top != -1){
        int tmp = obj->output.Data[obj->output.Top];
        return tmp;
    } else {
        while (obj->input.Top != -1){
            obj->output.Data[++obj->output.Top] = obj->input.Data[obj->input.Top--];
        }

        return obj->output.Data[obj->output.Top];
    }
}

// 两个栈都是空的则为空
bool myQueueEmpty(MyQueue* obj) {
    if (obj->input.Top == -1 && obj->output.Top == -1){
        return true;
    }
    return false;
}

// 释放即可
void myQueueFree(MyQueue* obj) {
    free(obj);
}