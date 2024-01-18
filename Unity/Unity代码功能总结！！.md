# Unity代码功能总结

#### 1.点击按钮缩放

```C#
private void OnMouseDown()
{
    transform.localScale = transform.localScale * 0.8f;
}

private void OnMouseUp()
{
    transform.localScale = transform.localScale / 0.8f;
}
```

#### 2.切换场景

```C#
using UnityEngine.SceneManagement;
SceneManager.LoadScene("xxx");
```

#### 3.相机跟随

```C#
public class CamerFollow1 : MonoBehaviour
{
    //跟随的游戏人物
    private Transform Player;
    //相机与人物之间的距离
    Vector3 offset;
    // Start is called before the first frame update
    void Start()
    {
        Player = GameObject.Find("JohnLemon").transform;
        offset = transform.position - Player.position;
    }

    // Update is called once per frame
    void Update()
    {
        transform.position = offset + Player.position;
    }
}
```

#### 4.鼠标点击跟随

```C#
private void Update()
{
        transform.position = Camera.main.ScreenToWorldPoint(Input.mousePosition);
        transform.position -= new Vector3(0,0,Camera.main.transform.position.z);
}
```

#### 5.获取当前物体名称

```C#
gameObject.name
```

#### 6.游戏对象旋转

```c#
//关闭z轴旋转
private void Start()
{
    //开启旋转
    rg.freezeRotation = false;
}
rg.freezeRotation = false;
```

#### 7.2D游戏角色移动

```C#
public class Player : MonoBehaviour
{
    Rigidbody2D rb;
    public float moveSpeed = 5f;
    private float moveX;
    void Start()
    {
        rb = GetComponent<Rigidbody2D>();
    }

    // Update is called once per frame
    void Update()
    {
        moveX = Input.GetAxisRaw("Horizontal");
        Move();
    }

    private void Move()    //移动函数
    {
        rb.velocity = new Vector2(moveX * moveSpeed, rb.velocity.y);
        if (moveX != 0)
        {
            transform.localScale = new Vector3(moveX, 1, 1);    //旋转
        }
    }
}
```

#### 8.2D游戏角色跳跃（带优化）

```C#
using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class Player : MonoBehaviour
{
    Rigidbody2D rb;
    [Range(1, 10)]
    public float jumpSpeed = 5f; //跳跃速度
    public bool isGrounded; //在地面
    public Transform groundCheck;   //地面检测点
    public LayerMask ground; //图层
    public float fallAdditon = 3.5f;//下落重力加成
    public float jumpAdditon = 1.5f;//下落重力加成
    public int jumpCount = 2; //多段跳次数
    private bool moveJump;  //跳跃输入
    private bool jumpHold; // 长按跳跃
    private bool isJump; //跳跃状态
    void Start()
    {
        rb = GetComponent<Rigidbody2D>();
        groundCheck = transform.Find("GroundCheck"); // 在角色底部创建一个用于检测地面的空物体
    }

    // Update is called once per frame
    void Update()
    {
        moveJump = Input.GetButtonDown("Jump");
        jumpHold = Input.GetButton("Jump");
        if (moveJump && jumpCount > 0)
        {
            isJump = true;
        }
    }
    private void FixedUpdate()
    {
        // 检测角色是否在地面上
        isGrounded = Physics2D.OverlapCircle(groundCheck.position, 0.1f, ground);
        Jump();
    }
    private void Jump()     //跳跃函数
    {
        if (isGrounded)		//复原多段跳
        {
            jumpCount = 2;
        }
        if (isJump)			//多段跳
        {
            rb.AddForce(Vector2.up * jumpSpeed, ForceMode2D.Impulse);
            jumpCount--;
            isJump = false;
        }
        if (rb.velocity.y < 0)
        {
            rb.gravityScale = fallAdditon;  //优化下落速度
        }
        else if (rb.velocity.y > 0 && !jumpHold)
        {
            rb.gravityScale = jumpAdditon;  //点按重力增加，区别长按
        }
        else
        {
            rb.gravityScale = 1;	//复原重力
        }
    }
}

```

#### 9.动画设置布尔值C

```C#
public void Run()
{
    bool playerHasXAisSpeed = Mathf.Abs(myRigidbody.velocity.x) > Mathf.Epsilon;    //检测速度
    myAnim.SetBool("isRun", playerHasXAisSpeed);    //动画设置布尔值
}
```

#### 10.角色攻击实现

```c#
public class PlayerAttack : MonoBehaviour
{
    public int damage = 2;
    public float startTime = 0.25f; //前摇时间
    public float endTime = 0.1f;    //后摇时间
    private Animator anim;			//角色动画
    private PolygonCollider2D coll2D;  //武器碰撞器
    void Start()
    {
        anim = GameObject.FindGameObjectWithTag("Player").GetComponent<Animator>();
        coll2D = GetComponent<PolygonCollider2D>();
    }

    // Update is called once per frame
    void Update()
    {
        Attack();
    }

    public void Attack()
    {
        if (Input.GetButtonDown("Attack"))
        {

            anim.SetTrigger("isAttack");
            StartCoroutine(StartAttack());    //启动协程，播放前摇
        }
    }
    IEnumerator StartAttack()
    {
        yield return new WaitForSeconds(startTime);		
        coll2D.enabled = true;
        StartCoroutine(DisableHitBox());	//启动协程，播放后摇，并关闭武器触发器
    }


    IEnumerator DisableHitBox()
    {
        yield return new WaitForSeconds(endTime);
        coll2D.enabled = false;			//关闭武器触发器
    }

    private void OnTriggerEnter2D(Collider2D collision)
    {
        if (collision.gameObject.CompareTag("Enemy"))
        {
            collision.GetComponent<Enemy>().TakeDamage(damage);		//承伤
        }
    }
}

```

#### 11.敌人受伤红色闪烁

```
public class Enemy : MonoBehaviour
{
    public int health = 5;  //血量
    public int hitDamage;   //攻击力
    public float flashTime = 0.2f;    //闪烁时间
    private SpriteRenderer sr;
    private Color originalColor;
    protected void Start()
    {
        sr = GetComponent<SpriteRenderer>();   
        originalColor = sr.color;   //暂存原始颜色
    }
    protected void Update()
    {
        if (health <= 0)
        {
            Destroy(gameObject);
        }
    }

    public void TakeDamage(int damage)  //受伤
    {
        health -= damage;
        FlashColor(flashTime);
    }

    public void FlashColor(float time)  //受伤红色闪烁
    {
        sr.color = Color.red;
        Invoke("ResetColor", time);
    }

    public void ResetColor()    //恢复原状
    {
        sr.color = originalColor;
    }
}

```

#### 12.敌人移动AI

```C#
public class EnemyBat : Enemy
{
    public float speed = 2; //移动速度
    public float startWaitTime = 1.5f;
    private float waitTime;//停留时间
    public Transform movePos;   //下次一次要移动的位置
    public Transform leftDownPos;  //移动的范围大小
    public Transform rightUpPos;   //移动的范围大小
    public void Start()
    {
        base.Start();
        waitTime = startWaitTime;
        movePos.position = GetRandomPos();
    }

    public void Update()
    {
        base.Update();
        transform.position = Vector2.MoveTowards(transform.position, movePos.position, speed * Time.deltaTime);
        if (Vector2.Distance(transform.position, movePos.position) < 0.1f)
        {
            if (waitTime <= 0)
            {
                movePos.position = GetRandomPos();
                waitTime = startWaitTime;
            }
            else
            {
                waitTime -= Time.deltaTime;
            }
        }
    }

    private Vector2 GetRandomPos()
    {
        Vector2 rndPos = new Vector2(Random.Range(leftDownPos.position.x, rightUpPos.position.x), Random.Range(leftDownPos.position.y, rightUpPos.position.y));
        return rndPos;
    }
}

```

#### 13.相机跟随2D

```C#
//主相机需要作为挂载这个脚本物体的子物体
public class CameraController : MonoBehaviour	
{
    public Transform target;	//跟随目标
    public float smoothing = 0.1f;	//平滑度

    void LateUpdate()
    {
        if (target != null)
        {
            if (transform.position != target.position)
            {
                Vector3 targetPos = target.position;
                transform.position = Vector3.Lerp(transform.position, targetPos, smoothing);	//向量插值
            }
        }
    }

}
```

#### 14.角色受伤闪烁

```c#
public class PlayerHealth : MonoBehaviour
{
    public int health;  //血量
    public int Blinks = 2;  //闪烁次数
    public float blinkTime = 0.1f;  //闪烁时间
    private Renderer myRender;

    // Start is called before the first frame update
    void Start()
    {
        myRender = GetComponent<Renderer>();
    }

    public void DamagePlayer(int damage)
    {
        health -= damage;
        if (health <= 0)
        {
            Destroy(gameObject);
        }
        BlinkPlayer(Blinks,blinkTime);
    }

    public void BlinkPlayer(int numBlinks, float seconds)
    {
        StartCoroutine(DoBlinks(numBlinks, seconds));
    }

    IEnumerator DoBlinks(int numBlinks, float seconds)
    {
        for (int i = 0; i < numBlinks * 2; i++)
        {
            myRender.enabled = !myRender.enabled;
            yield return new WaitForSeconds(seconds);
        }
        myRender.enabled = true;
    }
}

```

#### 15.角色血条

```c#
//需要将Image的图像类型替换成填充
public class HealthBar : MonoBehaviour
{
    public Text healthText;
    public static int HealthCurrent;
    public static int HealthMax=5;
    private Image healthBar;
    void Start()
    {
        healthBar = GetComponent<Image>();
        HealthCurrent = HealthMax;
    }

    void Update()
    {
        healthBar.fillAmount = (float)HealthCurrent / (float)HealthMax;
        healthText.text = HealthCurrent.ToString() + "/" + HealthMax.ToString();
    }
}

```

