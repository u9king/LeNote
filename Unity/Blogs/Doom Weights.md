---
title: Doom Weights
date: 2024-08-21 18:44:34
sticky: 13
pic: /images/thumb/Doom Weights.png
comments: true
toc: true
categories:
- Games
tags:
- Unity
---

# Doom Weights

#### 1. 游戏背景

​	《Doom Weights》是一款结合物理学与抽卡的回合制派对游戏。本作品是基于纸面模型设计，脱胎于TIG独立游戏工会与独立之光共同举办的纸上游戏工坊，作者将线下作品实现为线上游戏demo。引荐了小丑牌的卡牌效果设计。

#### 2. 团队成员

程序：u9king

#### 3. 游戏内容

<img src="/images/Unity/Games/Doom Weights/1.jpg">

<img src="/images/Unity/Games/Doom Weights/2.jpg">

<img src="/images/Unity/Games/Doom Weights/3.jpg">

#### 4. 核心机制
​	游戏需要三名玩家参与，胜利条件为最后存活的玩家获胜。游戏中存在很多负向卡牌来阻止玩家生存到下一个回合。游戏感受源于冰汽时代，希望带给玩家一种末日生存的体验。

- **演示视频**

<iframe src="//player.bilibili.com/player.html?bvid=BV1d8WAepEEJ" width="580px" height="320px"></iframe>

#### 5.核心代码

游戏卡牌的部分借鉴了很多小丑牌中的优秀效果

```C#
using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using UnityEngine.UI;
using TMPro;
using UnityEngine.EventSystems;
using UnityEngine.Events;

public class CardShow : MonoBehaviour, IDragHandler, IBeginDragHandler, IEndDragHandler, IPointerClickHandler
{
    [Header("数据挂载")]
    public CardSO myCard;
    public TMP_Text CardName;
    public Image BoarderBG;
    public TMP_Text FlavorText;
    public GameObject CardCover;
    public GameObject CardValue;
    public string CoverStr;

    [Header("卡牌移动")]
    public bool selected;
    public float selectionOffset = 50;
    private float lastClickTime = 0f;
    public float doubleClickThreshold = 0.3f;
    [HideInInspector] public UnityEvent<CardShow> BeginDragEvent;
    [HideInInspector] public UnityEvent<CardShow> EndDragEvent;
    private Image imageComponent;
    private RectTransform rectTransform;
    private Vector2 offset;


    private void Awake()
    {
        if(myCard != null)
        {
            CardName.text = myCard.cardName;
            BoarderBG.sprite = myCard.cardArt;
            FlavorText.text = myCard.flavorstr;
        }
        CardValue.SetActive(false); 
        CardCover.SetActive(true);
        //CardCover.GetComponentInChildren<TMP_Text>().text = CoverStr;
        imageComponent = GetComponent<Image>();
        rectTransform = GetComponent<RectTransform>();
    }

    #region 卡牌移动
    public void OnBeginDrag(PointerEventData eventData)
    {
        BeginDragEvent.Invoke(this);

        // 获取当前指针的位置并计算偏移量
        Vector2 localPointerPosition;
        RectTransformUtility.ScreenPointToLocalPointInRectangle(
            (RectTransform)transform.parent, eventData.position, eventData.enterEventCamera, out localPointerPosition);
        offset = localPointerPosition - rectTransform.anchoredPosition;

        imageComponent.raycastTarget = false; // 禁用射线检测，以便拖拽时不干扰其他UI
    }

    public void OnDrag(PointerEventData eventData)
    {
        // 更新卡片的位置
        Vector2 localPointerPosition;
        RectTransformUtility.ScreenPointToLocalPointInRectangle(
            (RectTransform)transform.parent, eventData.position, eventData.enterEventCamera, out localPointerPosition);
        rectTransform.anchoredPosition = localPointerPosition - offset;
    }

    public void OnEndDrag(PointerEventData eventData)
    {
        EndDragEvent.Invoke (this);
        imageComponent.raycastTarget = true; // 重新启用射线检测
        transform.localPosition = Vector3.zero;
    }

    public void OnPointerClick(PointerEventData eventData)
    {
        if (Time.time - lastClickTime < doubleClickThreshold && !CardController.isSelected)
        {
            CardController.isSelected = true;
            CardValue.SetActive(true);
            CardCover.SetActive(false);
            GameManager.instance.currentCard = myCard;
            GameManager.instance.CalculateCard(myCard);
            GameManager.instance.cardClicked = true;
            Invoke("CloseCardPanel", 1f);
            Debug.Log("Card double-clicked!");
        }
        lastClickTime = Time.time;
    }

    public void CloseCardPanel()
    {
        UIManager.instance.card_Panel.SetActive(false);
    }

    public int ParentIndex()
    {
        return transform.parent.GetSiblingIndex();
    }
    #endregion

    private void OnValidate()
    {
        if (CardCover != null)
        {
            TMP_Text coverText = CardCover.GetComponentInChildren<TMP_Text>();
            if (coverText != null)
            {
                coverText.text = CoverStr;
            }
        }
    }
}

```

#### 6.优化方向


- 增加卡牌类型
- 优化数值设计
- 增加多种不同重量的砝码
- 增加四角回弹机制
- 优化多人数参与可能性
- 增加卡牌特效
- 线上联机功能
- 增加砝码的三维堆叠效果
- 增加背景故事