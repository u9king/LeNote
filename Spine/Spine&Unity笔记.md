# Spine&Unity笔记

#### 1.

```C#
using Spine;
using Spine.Unity;
using System.Collections;
using UnityEngine;

public class SpineAnim : MonoBehaviour
{
    [SpineAnimation]
    public string attackAnimationName;
    [SpineAnimation]
    public string idleAnimationName;
    [SpineAnimation]
    public string walkAnimationName;
    [SpineEvent(dataField: "skeletonAnimation", fallbackToTextField: true)]
    public string flystepEventName;

    [Header("监视")]
    public SkeletonAnimation skeletonAnimation;
    //public Spine.AnimationState spineAnimationState;
    public Spine.Skeleton skeleton;
    public Spine.EventData eventData;


    IEnumerator Start()
    {
        skeletonAnimation = GetComponent<SkeletonAnimation>();  //动画
        //spineAnimationState = skeletonAnimation.AnimationState;  //状态机
        skeleton = skeletonAnimation.skeleton;  //骨骼
        skeleton.ScaleX = -1f;
        eventData = skeleton.Data.FindEvent(flystepEventName);  //事件
        skeletonAnimation.AnimationState.Event += HandleAnimationStateEvent;
        if (skeletonAnimation == null) yield break;
        while (true)
        {
            //Debug.Log(1);
            skeletonAnimation.AnimationState.SetAnimation(1, walkAnimationName, false);
            skeletonAnimation.AnimationState.AddAnimation(1, idleAnimationName, true,0);
            skeleton.ScaleX *= -1f;
            yield return new WaitForSeconds(Random.Range(minimumDelay, maximumDelay));
        }
    }

    private void HandleAnimationStateEvent(TrackEntry trackEntry, Spine.Event e)
    {
        bool eventMatch = (eventData == e.Data);
        if (eventMatch)
        {
            audioSource.pitch = Random.Range(0.2f, 2f);
            audioSource.clip = audioClip;
            audioSource.Play();
        }
    }

    [Header("其他设置")]
    public float minimumDelay = 1f;
    public float maximumDelay = 3f;
    public AudioSource audioSource;
    public AudioClip audioClip;
}
```

